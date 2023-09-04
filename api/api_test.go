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
		FirstPrice: 30787.2, LastPrice: 30767.9, HightPrice: 30980.5, LowPrice: 30761.8,

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
	fmt.Println(k.MultipleOfTheKKind(k.Kind(30980, 30485), 30980, 30485.0, 2))

}

func TestApi1(t *testing.T) {
	fmt.Println(47 / (31145.0 - 30600.0))
}
