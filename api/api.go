package api

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/client/http"
	"github.com/liujunren93/bian/client/websocket"
)

const (
	ENUM_SIDE_BUY  = "buy"
	ENUM_SIDE_SELL = "sell"
	SUB_BASE_API   = "wss://fstream.binance.com/ws"
)

type Api struct {
	ApiKey    string
	ApiSecret string
	BaseApi   string
}

func (a *Api) HttpClient(baseApi string) *http.HttpClient {
	return http.NewClient(client.Config{
		BaseApi:   baseApi,
		ApiKey:    a.ApiKey,
		ApiSecret: a.ApiSecret,
	})
}

func (a *Api) WsClient(baseApi string) *websocket.WsClient {
	return websocket.NewClient(client.Config{
		BaseApi:   baseApi,
		ApiKey:    a.ApiKey,
		ApiSecret: a.ApiSecret,
	})
}

type Interval string

func (i Interval) String() string {
	return string(i)
}

// func (i Interval) ToDuration() (time.Duration, string) {
// 	return i.toDuration()

// }

func (i Interval) Unit() string {
	_, u := i.ToDuration()
	return u

}

func (i Interval) GetInterval() float32 {
	_, unit := i.ToDuration()
	switch unit {
	case "m":
		return 0.01
	case "h":
		return 0.1
	case "d":
		return 0.2
	case "w":
		return 0.1
	case "M":
		return 0.1
	}
	return 0.1

}

func (i Interval) ToDuration() (time.Duration, string) {
	matched := regexp.MustCompile("[A-Za-z]$")
	str := string(i)
	sl := matched.FindStringIndex(str)
	unit := str[sl[0]:sl[1]]
	num, _ := strconv.Atoi(str[:sl[0]])
	switch unit {
	case "m":
		return time.Duration(num) * time.Minute, "m"
	case "h":
		return time.Duration(num) * time.Hour, "h"
	case "d":
		return time.Duration(num) * time.Hour * 24, "d"
	case "w":
		return time.Duration(num) * time.Hour * 24 * 7, "w"
	case "M":
		return time.Duration(num) * time.Hour * 24 * 30, "M"
	}
	return 0, str
}

type KLine struct {
	Symbol     string   `json:"s"`
	BeginTime  int64    `json:"t"` // 这根K线的起始时间
	EndTime    int64    `json:"T"` // 这根K线的结束时间
	Interval   Interval `json:"i"` // K线间隔
	FirestID   int      `json:"f"` // 这根K线期间第一笔成交ID
	LastID     int      `json:"L"` // 这根K线期间末一笔成交ID
	FirstPrice float64  `json:"o"` // 这根K线期间第一笔成交价
	LastPrice  float64  `json:"c"` // 这根K线期间末一笔成交价
	HightPrice float64  `json:"h"` // 这根K线期间最高成交价
	LowPrice   float64  `json:"l"` // 这根K线期间最低成交价
	Volume     float64  `json:"v"` // 这根K线期间成交量
	Cnt        int      `json:"n"` // 这根K线期间成交笔数
	IsOver     bool     `json:"x"` // 这根K线是否完结
	Amount     float64  `json:"q"` // 这根K线期间成交额
	V          float64  `json:"V"` // 主动买入的成交量
	Q          float64  `json:"Q"` // 主动买入的成交额
}

type KLineKind int8

func (k KLineKind) String() string {
	switch k {

	case KLineKind_YANG_1:
		return "大阳线"
	case KLineKind_YANG_2:
		return "中阳线"
	case KLineKind_YANG_3:
		return "小阳线"
	case KLineKind_YANG_4:
		return "阳十字"
	case KLineKind_YIN_1:
		return "大阴线"
	case KLineKind_YIN_2:
		return "中阴线"
	case KLineKind_YIN_3:
		return "小阴线"
	case KLineKind_YIN_4:
		return "阴十字"
	case KLineKind_YANG_PINBAR:
		return "阳pinbar"
	case KLineKind_YIN_PINBAR:
		return "阴pinbar"
	default:
		return "未确定"
	}
}

const (
	KLineKind_UNKNOWN KLineKind = iota // 未close
	KLineKind_YANG_1                   // 大阳线
	KLineKind_YANG_2                   // 中阳线
	KLineKind_YANG_3                   // 小阳线
	KLineKind_YANG_4                   // 阳十字
	KLineKind_YANG_PINBAR
	KLineKind_YIN_1 // 大阴线
	KLineKind_YIN_2 // 中阴线
	KLineKind_YIN_3 // 小阴线
	KLineKind_YIN_4 // 阴十字
	KLineKind_YIN_PINBAR
)

func (k *KLine) Kind(level []float32) KLineKind {
	if !k.IsOver {
		return KLineKind_UNKNOWN
	}
	const day = time.Hour * 24
	// t, _ := k.Interval.toDuration()
	// times := float64(day) / float64(t)
	if ((k.HightPrice-k.LowPrice)-math.Abs(k.FirstPrice-k.LastPrice))/math.Abs(k.FirstPrice-k.LastPrice) >= 2 {
		if k.LastPrice > k.FirstPrice {
			return KLineKind_YANG_PINBAR
		} else {
			return KLineKind_YIN_PINBAR
		}
	}
	switch {
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(level[0]): //大阳
		return KLineKind_YANG_1

	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(level[1]): //中阳
		return KLineKind_YANG_2
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(level[2]): //小阳
		return KLineKind_YANG_3
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= 0: //阳十字
		return KLineKind_YANG_4
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(level[0]): //大阴
		return KLineKind_YIN_1
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(level[1]): //中阴
		return KLineKind_YIN_2
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(level[2]): //小阴
		return KLineKind_YIN_3
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= 0: //阴十字
		return KLineKind_YIN_4

	}

	return KLineKind_UNKNOWN

}
