package api

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	var k = KLine{
		Symbol:     "",
		BeginTime:  0,
		EndTime:    0,
		Interval:   "15m",
		FirestID:   0,
		LastID:     0,
		FirstPrice: 26946.80,
		HightPrice: 26946.80,
		LowPrice:   26830,
		LastPrice:  26893,

		Volume: 0,
		Cnt:    0,
		IsOver: true,
		Amount: 0,
		V:      0,
		Q:      0,
	}
	// - 0.002
	//         - 0.003
	//         - 0.005
	// 0.0015 0.0025 0.0035
	// 0.0011
	fmt.Println(k.Kind([]float32{0.005, 0.0003, 0.0002}))

}
