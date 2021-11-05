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
	var maxProduct int64 = 0
OUTER:
	for idx := 0; idx <= len(digits)-span; idx++ {
		if digits[idx] == '0' {
			continue
		}
		var product int64 = int64(digits[idx] - '0')
		for i := 1; i < span; i++ {
			if digits[idx+i] < '0' || digits[idx+i] > '9' {
				return 0, errors.New("digits input must only contain digits")
			} else if digits[idx+i] == '0' {
				continue OUTER
			}
			product *= int64(digits[idx+i] - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}
