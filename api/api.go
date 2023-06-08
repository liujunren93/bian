package api

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/client/http"
	"github.com/liujunren93/bian/client/websocket"
)

type SideType int8

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
func (i Interval) Name() string {
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

func (k *KLine) String() string {
	return fmt.Sprintf("ID:%d,{FirstPrice:%v,LastPrice:%v,HightPrice:%v,LowPrice:%v},Amplitude:%v", k.FirestID, k.FirstPrice, k.LastPrice, k.HightPrice, k.LowPrice, k.Amplitude())
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

func (k *KLine) IsStatic(levels []float32) bool {
	kind := k.kind(levels)
	return kind == KLineKind_YANG_4 || kind == KLineKind_YANG_PINBAR || kind == KLineKind_YIN_4 || kind == KLineKind_YIN_PINBAR

}
func (k *KLine) Kind(levels []float32) KLineKind {
	if !k.IsOver {
		return KLineKind_UNKNOWN
	}
	// t, _ := k.Interval.toDuration()
	return k.kind(levels)

}

func (k *KLine) KindIgnoreOver(levels []float32) KLineKind {

	// t, _ := k.Interval.toDuration()
	return k.kind(levels)

}

func (k *KLine) MultipleOfTheKKind(kd KLineKind, levels []float32) float64 {

	if kd == KLineKind_YANG_1 || kd == KLineKind_YIN_1 {
		return math.Abs((k.LastPrice-k.FirstPrice)/k.FirstPrice) / float64(levels[0])
	}
	if kd == KLineKind_YANG_2 || kd == KLineKind_YIN_2 {
		return math.Abs((k.LastPrice-k.FirstPrice)/k.FirstPrice) / float64(levels[1])
	}
	if kd == KLineKind_YANG_3 || kd == KLineKind_YIN_3 {
		return math.Abs((k.LastPrice-k.FirstPrice)/k.FirstPrice) / float64(levels[2])
	}
	return 0

}

func (k *KLine) kind(levels []float32) (klkind KLineKind) {
	const day = time.Hour * 24
	// t, _ := k.Interval.toDuration()
	// times := float64(day) / float64(t)
	switch {
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(levels[0]): //大阳
		klkind = KLineKind_YANG_1

	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(levels[1]): //中阳
		klkind = KLineKind_YANG_2
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= float64(levels[2]): //小阳
		klkind = KLineKind_YANG_3
	case (k.LastPrice-k.FirstPrice)/k.FirstPrice >= 0: //阳十字
		klkind = KLineKind_YANG_4
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(levels[0]): //大阴
		klkind = KLineKind_YIN_1
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(levels[1]): //中阴
		klkind = KLineKind_YIN_2
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= float64(levels[2]): //小阴
		klkind = KLineKind_YIN_3
	case (k.FirstPrice-k.LastPrice)/k.FirstPrice >= 0: //阴十字
		klkind = KLineKind_YIN_4

	}
	if klkind == KLineKind_YIN_4 || klkind == KLineKind_YIN_3 || klkind == KLineKind_YANG_4 || klkind == KLineKind_YANG_3 {
		if ((k.HightPrice-k.LowPrice)-math.Abs(k.FirstPrice-k.LastPrice))/math.Abs(k.FirstPrice-k.LastPrice) >= 2 {
			if k.LastPrice > k.FirstPrice {
				klkind = KLineKind_YANG_PINBAR
			} else {
				klkind = KLineKind_YIN_PINBAR
			}
		}
	}

	return

}

type LeadLevel int8

const (
	LEAD_LEVEL_SMALL LeadLevel = iota
	LEAD_LEVEL_MEDIUM
	LEAD_LEVEL_LARGE
)

// 振幅 实体长度<0:阴线 >0:阳线
func (k *KLine) Amplitude() float64 {
	return k.LastPrice - k.FirstPrice
}

func (k *KLine) BigAmplitude() float64 {
	return k.HightPrice - k.LowPrice
}

func (k *KLine) LeadLevel(levels []float32) (top, down LeadLevel) {
	top, down = LEAD_LEVEL_LARGE, LEAD_LEVEL_LARGE
	var leadChangeTop float64
	var leadChangeDown float64
	kind := k.kind(levels)
	if k.Amplitude() > 0 {
		leadChangeTop = k.HightPrice - k.LastPrice
		leadChangeDown = k.FirstPrice - k.LowPrice
	} else {
		leadChangeTop = k.HightPrice - k.FirstPrice
		leadChangeDown = k.LastPrice - k.LowPrice
	}
	amplitude := math.Abs(k.Amplitude())
	switch {
	case kind == KLineKind_YANG_1 || kind == KLineKind_YIN_1:
		if leadChangeTop >= amplitude {
			top = LEAD_LEVEL_LARGE
		} else if leadChangeTop >= amplitude/2 {
			top = LEAD_LEVEL_MEDIUM
		}
		if leadChangeDown > amplitude {
			down = LEAD_LEVEL_LARGE
		}
		if leadChangeDown > amplitude/2 {
			down = LEAD_LEVEL_MEDIUM
		}
	case kind == KLineKind_YANG_2 || kind == KLineKind_YIN_2:
		if leadChangeTop >= amplitude*1.5 {
			top = LEAD_LEVEL_LARGE
		} else if leadChangeTop >= amplitude {
			top = LEAD_LEVEL_MEDIUM
		}
		if leadChangeDown > amplitude*1.5 {
			down = LEAD_LEVEL_LARGE
		} else if leadChangeDown >= amplitude {
			down = LEAD_LEVEL_MEDIUM
		}
	case kind == KLineKind_YANG_3 || kind == KLineKind_YIN_3:
		if leadChangeTop >= amplitude*3 {
			top = LEAD_LEVEL_LARGE
		} else if leadChangeTop >= amplitude*1.5 {
			top = LEAD_LEVEL_MEDIUM
		}
		if leadChangeDown > amplitude*3 {
			down = LEAD_LEVEL_LARGE
		} else if leadChangeDown >= amplitude*1.5 {
			down = LEAD_LEVEL_MEDIUM
		}
	case kind == KLineKind_YANG_4 || kind == KLineKind_YIN_4 || kind == KLineKind_YANG_PINBAR || kind == KLineKind_YIN_PINBAR:
		if leadChangeTop >= amplitude*4 {
			top = LEAD_LEVEL_LARGE
		} else if leadChangeTop >= amplitude*2 {
			top = LEAD_LEVEL_MEDIUM
		}
		if leadChangeDown > amplitude*4 {
			down = LEAD_LEVEL_LARGE
		} else if leadChangeDown >= amplitude*2 {
			down = LEAD_LEVEL_MEDIUM
		}

	}

	return

}
