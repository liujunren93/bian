package spot

import (
	"fmt"
	"strings"

	"github.com/liujunren93/bian/api"
	"github.com/liujunren93/bian/client"
)

type spot struct {
	api.Api
}

func NewSpot(conf *client.Config) *spot {
	return &spot{
		Api: api.Api{
			ApiKey:    conf.ApiKey,
			ApiSecret: conf.ApiSecret,
			BaseApi:   baseApi,
		},
	}
}

const baseApi = "https://api.binance.com"

func (f *spot) TradeParams(symbol, side string, params client.Params) error {
	var res map[string]interface{}

	params["timeInForce"] = "GTC"
	params["symbol"] = strings.ToUpper(symbol)
	params["side"] = strings.ToUpper(side)

	err := f.HttpClient(baseApi).Post("/api/v3/order", nil, params, nil, &res)

	if res["code"] != 200 {

		return fmt.Errorf("%v", res)
	}
	fmt.Println(err)
	return err
}
