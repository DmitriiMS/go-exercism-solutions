// Package luhn provides function that validates ID numbers with Luhn algorithm
package luhn

import (
	"strings"
	"unicode"
)

// Valid chackes if given ID number is valid by using the Luhn algorithm.
func Valid(susNumber string) bool {
	// Trim spaces left and right, split string by spaces and glue back together without separator.
	trimmedRunes := []rune(strings.Join(strings.Split(strings.TrimSpace(susNumber), " "), ""))
	if len(trimmedRunes) <= 1 { // simple check for minimal length
		return false
	}
	// Walk over the prepared rune slice.
	// If something other than digit is encountered, return false.
	// If everything is ok, proceed with teh algorithm.
	sum := 0
	for marker, i := len(trimmedRunes)-2, len(trimmedRunes)-1; i >= 0; i-- {
		if unicode.IsDigit(trimmedRunes[i]) {
			digit := int(trimmedRunes[i] - '0')
			if i == marker {
				marker -= 2
				if digit*2 > 9 {
					sum += digit*2 - 9
					continue
				}
				sum += digit * 2
				continue
			}
			sum += digit
			continue
		}
		return false
	}
	return sum%10 == 0 // check if ID is valid
}
