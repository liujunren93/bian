package future

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/liujunren93/bian/api"
	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/client/websocket"
	"github.com/mitchellh/mapstructure"
)

type Future struct {
	api.Api
}

const baseApi = "https://fapi.binance.com"

func NewFuture(conf *client.Config) *Future {
	return &Future{
		Api: api.Api{
			ApiKey:    conf.ApiKey,
			ApiSecret: conf.ApiSecret,
			BaseApi:   baseApi,
		},
	}
}

func (f *Future) TradeParams(symbol, side string, params client.Params) error {
	side = strings.ToUpper(side)
	var res map[string]interface{}
	if side != "" {
		params["side"] = side
	}

	params["symbol"] = strings.ToUpper(symbol)

	err := f.HttpClient(baseApi).Post("/fapi/v1/order", nil, params, nil, &res)

	if _, ok := res["code"]; ok && res["code"] != 200 {
		return fmt.Errorf("res:%v,params:%#v", res, params)
	}
	return err
}
func (f *Future) CancelOrder(symbol string, params client.Params) error {

	var res map[string]interface{}

	params["symbol"] = strings.ToUpper(symbol)

	err := f.HttpClient(baseApi).Delete("/fapi/v1/order", nil, params, nil, &res)

	if _, ok := res["code"]; ok && res["code"] != 200 {
		return fmt.Errorf("res:%v,params:%#v", res, params)
	}
	return err
}

func (f *Future) CancelAllOrder(symbol string) error {

	var res map[string]interface{}
	var params = client.Params{
		"symbol": strings.ToUpper(symbol),
	}
	params["symbol"] = strings.ToUpper(symbol)
	h := http.Header{}
	h.Add("Content-Type", "application/x-www-form-urlencoded")
	err := f.HttpClient(baseApi).Delete("/fapi/v1/allOpenOrders", h, params, nil, &res)

	if _, ok := res["code"]; ok && res["code"] != 200 {
		return fmt.Errorf("res:%v,params:%#v", res, params)
	}
	return err
}

func (f *Future) SubscribeKline(path string, params []string, callback func(*api.KLine, error)) error {
	isOne := true
	baseApi := "wss://fstream.binance.com/ws"
	// if path == "" {
	// 	isOne = false
	// 	baseApi = "wss://fstream.binance.com/stream/?"
	// 	for i, p := range params {
	// 		if i != 0 {
	// 			path += "/"
	// 		}
	// 		path += p
	// 	}
	// }
	// fmt.Println(path)
	cli := f.WsClient(baseApi)

	go func() {
	RECONNECT:
		cli.Send(path, nil, websocket.Params{
			ID:     time.Now().Unix(),
			Method: "SUBSCRIBE",
			Params: params,
		})

		done, err := cli.Receiver(path, func(b []byte) {
			var tmpKlmap map[string]interface{}
			var symbol string
			var kl api.KLine
			var res map[string]interface{}
			err := json.Unmarshal(b, &res)
			if err != nil {
				callback(nil, err)
				return
			}

			if isOne {

				if res["k"] != nil {
					tmpKlmap = res["k"].(map[string]interface{})
					symbol = res["s"].(string)
					// symbol = strings.Split(res["s"].(string), "_")[1]
				}
			} else {

				if res["data"] == nil {
					return
				}
				symbol = res["stream"].(string)
				if val, ok := res["data"].(map[string]interface{}); ok {
					if val["k"] != nil {
						if va, ok := val["k"].(map[string]interface{}); ok {
							tmpKlmap = va
						}
					}

				}
				// fmt.Println(res["data"].(map[string]interface{}))
			}
			if tmpKlmap == nil {
				return
			}
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &kl})
			if err != nil {
				callback(nil, err)
				return
			}

			decoder.Decode(tmpKlmap)
			if strings.Contains(symbol, "@") {
				symbol = strings.Split(symbol, "@")[0]
			}
			kl.FirestID = int(kl.BeginTime) / 1000
			kl.FirstPrice, _ = strconv.ParseFloat(tmpKlmap["o"].(string), 64)
			kl.LastPrice, _ = strconv.ParseFloat(tmpKlmap["c"].(string), 64)
			kl.HightPrice, _ = strconv.ParseFloat(tmpKlmap["h"].(string), 64)
			kl.LowPrice, _ = strconv.ParseFloat(tmpKlmap["l"].(string), 64)
			kl.Volume, _ = strconv.ParseFloat(tmpKlmap["v"].(string), 64)
			kl.Amount, _ = strconv.ParseFloat(tmpKlmap["q"].(string), 64)
			kl.V, _ = strconv.ParseFloat(tmpKlmap["V"].(string), 64)
			kl.Q, _ = strconv.ParseFloat(tmpKlmap["Q"].(string), 64)
			kl.Symbol = symbol
			kl.BeginTime = kl.BeginTime / 1000
			kl.EndTime = kl.EndTime / 1000
			callback(&kl, nil)

		})

		callback(nil, err)
		<-done
		goto RECONNECT
	}()
	return nil
}
