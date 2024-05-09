package spot

import "github.com/shopspring/decimal"

type Side string

type OrderType string

const (
	ORDER_TYPE_LIMIT             OrderType = "LIMIT"             // 限价单 [timeInForce, quantity, price]
	ORDER_TYPE_MARKET            OrderType = "MARKET"            //市价单 [quantity 或者 quoteOrderQty]
	ORDER_TYPE_STOP_LOSS         OrderType = "STOP_LOSS"         //止损单 [quantity, stopPrice 或者 trailingDelta ]
	ORDER_TYPE_STOP_LOSS_LIMIT   OrderType = "STOP_LOSS_LIMIT"   //限价止损单 [timeInForce, quantity, price, stopPrice 或者 trailingDelta]
	ORDER_TYPE_TAKE_PROFIT       OrderType = "TAKE_PROFIT "      //止盈单 [quantity, stopPrice 或者 trailingDelta]
	ORDER_TYPE_TAKE_PROFIT_LIMIT OrderType = "TAKE_PROFIT_LIMIT" //限价止盈单 [timeInForce, quantity, price, stopPrice 或者 trailingDelta]
	ORDER_TYPE_LIMIT_MAKER       OrderType = "LIMIT_MAKER"       //限价只挂单[quantity, price]
)

type TimeInForce string

const (
	TIME_IN_FORCE_GTC TimeInForce = "GTC" //成交为止
	TIME_IN_FORCE_IOC TimeInForce = "IOC" //无法立即成交的部分就撤销,订单在失效前会尽量多的成交。
	TIME_IN_FORCE_FOK TimeInForce = "FOK" //无法全部立即成交就撤销,如果无法全部成交，订单会失效。
)

type OrderTradeRequest struct {
	Symbol           string          `json:"symbol"` //required
	Side             string          `json:"side"`   //required //BUY /SELL
	Type             OrderType       `json:"type"`   //required
	TimeInForce      TimeInForce     `json:"timeInForce"`
	Quantity         decimal.Decimal `json:"quantity"` // 数量
	QuoteOrderQty    decimal.Decimal `json:"quoteOrderQty"`
	Price            decimal.Decimal `json:"price"`
	NewClientOrderId string          `json:"newClientOrderId"` //	客户自定义的唯一订单ID
	StopPrice        decimal.Decimal `json:"stopPrice"`        //仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT 和 TAKE_PROFIT_LIMIT 需要此参数
	TrailingDelta    int             `json:"trailingDelta"`
}

type OrderStatus string

const (
	ORDER_STATUS_NEW              OrderStatus = "NEW"              // 创建成功
	ORDER_STATUS_PARTIALLY_FILLED OrderStatus = "PARTIALLY_FILLED" //部分订单被成交
	ORDER_STATUS_FILLED           OrderStatus = "FILLED"           // 订单完全成交
	ORDER_STATUS_CANCELED         OrderStatus = "CANCELED"         //用户撤销了订单
	ORDER_STATUS_PENDING_CANCEL   OrderStatus = "PENDING_CANCEL"   //	撤销中（目前并未使用）
	ORDER_STATUS_REJECTED         OrderStatus = "REJECTED"         //订单没有被交易引擎接受，也没被处理
	ORDER_STATUS_EXPIRED          OrderStatus = "EXPIRED"          //订单被交易引擎取消，
	ORDER_STATUS_EXPIRED_IN_MATCH OrderStatus = "EXPIRED_IN_MATCH" //表示订单由于 STP 触发而过期
)

type OrderBaseResponse struct {
	Symbol                  string                 `json:"symbol"`
	OrderID                 string                 `json:"orderId"`
	ClientOrderId           string                 `json:"clientOrderId"`           //	客户自定义的唯一订单ID
	TransactTime            int                    `json:"transactTime"`            //交易的时间戳 毫秒
	Price                   decimal.Decimal        `json:"price"`                   // 订单价格
	OrigQty                 decimal.Decimal        `json:"origQty"`                 //用户设置的原始订单数量
	ExecutedQty             decimal.Decimal        `json:"executedQty"`             //交易的订单数量
	CummulativeQuoteQty     decimal.Decimal        `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Side                    string                 `json:"side"`                    // //BUY /SELL
	Type                    OrderType              `json:"type"`                    //订单类型
	TimeInForce             TimeInForce            `json:"timeInForce"`             //订单的时效方式
	WorkingTime             int                    `json:"workingTime"`             // 订单添加到 order book 的时间
	SelfTradePreventionMode string                 `json:"selfTradePreventionMode"` // 自我交易预防模式
	Status                  OrderStatus            `json:"status"`
	Fills                   OrderBaseResponseFills `json:"fills"`
}

type OrderBaseResponseFills struct {
	Price           string `json:"price"`           // 交易的价格
	Qty             string `json:"qty"`             // 交易的数量
	Commission      string `json:"commission"`      // 手续费金额
	CommissionAsset string `json:"commissionAsset"` // 手续费的币种
	TradeID         int    `json:"tradeId"`         // 交易ID
}
