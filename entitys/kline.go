package entitys

import "github.com/shopspring/decimal"

type KlineResponse struct {
	Event  string `json:"e"` // 事件类型
	Time   int64  `json:"E"` // 事件时间
	Symbol string `json:"s"` // 交易对
	Kine   Kine   `json:"k"`
}
type Kine struct {
	BeginTime             int             `json:"t"` //这根K线的起始时间
	EndTime               int             `json:"T"` // 这根K线的结束时间
	Symbol                string          `json:"s"`
	Interval              string          `json:"i"` // K线间隔
	FirstOrderID          int             `json:"f"` // 这根K线期间第一笔成交ID
	LastOrderID           int             `json:"L"` // 这根K线期间末一笔成交ID
	FirstTradePrice       decimal.Decimal `json:"o"` // 这根K线期间第一笔成交价
	LastTradePrice        decimal.Decimal `json:"c"` // 这根K线期间末一笔成交价
	HightTradePrice       decimal.Decimal `json:"h"` // 这根K线期间最高成交价
	LowTradePrice         decimal.Decimal `json:"l"` // 这根K线期间最低成交价
	TradeTotal            decimal.Decimal `json:"v"` // 这根K线期间成交量
	TradeCount            int             `json:"n"` // 这根K线期间成交笔数
	IsFinish              bool            `json:"x"` // 这根K线是否完结(是否已经开始下一根K线)
	TradeAmount           decimal.Decimal `json:"q"` // 这根K线期间成交额
	InitiativeTradeTotal  decimal.Decimal `json:"V"` // 主动买入的成交量
	InitiativeTradeAmount decimal.Decimal `json:"Q"` // 主动买入的成交额
}
