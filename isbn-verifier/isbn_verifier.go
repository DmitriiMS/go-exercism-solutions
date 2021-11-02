package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	if len(isbn) == 0 {
		return false
	}
	sum, i := 0, 10
	for _, c := range isbn {
		if c == 'X' && i == 1 {
			sum += 10
			i = 0
		}
		if unicode.IsDigit(c) {
			sum += int(c-'0') * i
			i--
		}
	}
	if i != 0 || sum%11 != 0 {
		return false
	}
	return true
}
