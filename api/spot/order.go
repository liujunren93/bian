package spot

import (
	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/client/http"
)

type order struct {
}

func NewOrder(key, secret string) *order {

}

func (o *order) GetAllOrder() []map[string]interface{} {
	var res []map[string]interface{}
	err := o.HttpClient().Get("/api/v3/openOrders", nil, client.QueryParams{"symbol": "BTCUSDT"}, &res)
	if err != nil {
		panic(err)
	}
	return res
}

func (o *order) Trade(side string, price, quantity float64) int64 {
	var res map[string]interface{}
	spotClient.Post("/api/v3/order", nil, client.Params{
		"symbol":      "BTCUSDT",
		"side":        side,
		"type":        "LIMIT",
		"price":       price,
		"quantity":    quantity,
		"timeInForce": "GTC",
	}, nil, &res)
	/* fmt.Println(res)
	fmt.Println("trade_res:", res["orderId"].(float64))*/

	/*	if res["code"] != 200 {
		log.Log.Error(res)
		return 0
	}*/
	return int64(res["orderId"].(float64))
}
