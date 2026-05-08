package system

import (
	"slices"
	"math"
)

// binary system
func BinToDec(bin []int) int {
	slices.Reverse(bin)
	
	var dec int = 0


	for index, value := range bin {
		dec += value * int(math.Pow(2, float64(index)))
	}


	return dec
}