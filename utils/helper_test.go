package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	var a = []int{1, 2, 3}
	a = append(a, 5)
	a = a[1:]
	fmt.Print(a)
}
