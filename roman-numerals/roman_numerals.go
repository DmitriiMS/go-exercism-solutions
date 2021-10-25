//packafe romannumerals provides function that converts arabic numerals to roman numerals
package romannumerals

import (
	"errors"
)

//compareSlice holds numerals that we compare incoming number against, this is needed for conputations
//numbersToRoman links arabic numerals to roman numerals
var compareSlice = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var arabicToRoman = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

//ToRomanNumeral takes arabic numeral and returns its roman representation. It also returns error if anything goes wrong duting the conversion.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 { //check if number isn't too small or too big
		return "", errors.New("u wot m8?")
	}
	result := ""
	for _, element := range compareSlice {
		for element <= arabic {
			arabic -= element
			result += arabicToRoman[element]
		}
	}
	return result, nil
}
