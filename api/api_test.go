package api

import (
	"fmt"
	"sync"
	"testing"
)

func TestApi(t *testing.T) {
	sync.NewCond(l sync.Locker)
	var k = KLine{
		Symbol:     "",
		BeginTime:  0,
		EndTime:    0,
		Interval:   "15m",
		FirestID:   0,
		LastID:     0,
		FirstPrice: 28071.8,
		HightPrice: 28084,
		LowPrice:   27651,
		LastPrice:  27698.5,

		Volume: 0,
		Cnt:    0,
		IsOver: true,
		Amount: 0,
		V:      0,
		Q:      0,
	}

	fmt.Println(k.Kind([]float32{0.005, 0.003, 0.001}))

}
