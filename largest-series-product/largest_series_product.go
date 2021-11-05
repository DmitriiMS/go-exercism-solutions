package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must not be negative")
	} else if span == 0 {
		return 1, nil
	}
	var product, maxProduct int64 = 1, 0
	subLen := 0
	for idx := 0; idx < len(digits); idx++ {
		if digits[idx] == '0' {
			product, subLen = 1, 0
			continue
		}
		if digits[idx] < '0' || digits[idx] > '9' {
			return 0, errors.New("digits input must only contain digits")
		}
		if subLen < span {
			product *= int64(digits[idx] - '0')
			subLen++
		}
		if subLen == span {
			if product > maxProduct {
				maxProduct = product
			}
			//divide by the leftmost number in series
			//after that we need to mutiply by the new number on the right to get the new product
			product /= int64(digits[idx-span+1] - '0')
			subLen--
		}
	}
	return maxProduct, nil
}
