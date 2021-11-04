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
	for start, end := 0, span; end <= len(digits); start, end = start+1, end+1 {
		product, err := calculateProduct([]rune(digits[start:end]))
		if err != nil {
			return -1, err
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}

func calculateProduct(str []rune) (int64, error) {
	var product int64 = 1
	for _, c := range str {
		if c < '0' || c > '9' {
			return 0, errors.New("digits input must only contain digits")
		}
		product *= int64(c - '0')
	}
	return product, nil
}
