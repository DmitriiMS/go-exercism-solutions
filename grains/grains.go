package grains

import (
	"errors"
	"math"
	"strconv"
)

// gogle time, lets find us some geometric progression formulas
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("invalid square number, valid square numbers are [1...64]")
	}
	return uint64(math.Pow(2, float64(square-1))), nil
}

func Total() uint64 {
	var temp string = strconv.FormatFloat(math.Pow(2, 64)-1, 'f', 0, 64)
	var ret uint64 = 0
	ret, _ = strconv.ParseUint(temp, 10, 64)
	return ret
}
