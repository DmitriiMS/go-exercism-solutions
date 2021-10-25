//packafe romannumerals provides function that converts arabic numerals to roman numerals
package romannumerals

import (
	"errors"
	"strings"
)

//compareSlice holds numerals that we compare incoming number against, this is needed for conputations
//numbersToRoman links arabic numerals to roman numerals
var compareSlice = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numbersToRoman = map[int]string{
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

type Roman string //custom type that allows creating a method on a string

//ToRomanNumeral takes arabic numeral and returns its roman representation. It also returns error if anything goes wrong duting the conversion.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 { //check if number isn't too small or too big
		return "", errors.New("u wot m8?")
	}
	var result Roman
	result.recursiveBuilder(arabic) //if everything is ok, call recursive method on result, which is empty at the moment
	return string(result), nil
}

//recursiveBuilder takes given arabic numeral and recursively calculates roman representation for each position in the number,
//starting with the leftmost, and concateneates result of each iteration to already existing string
//First, remainder is set to 0.
//Then, for each number in compareSlice, it checks if it's less or equal to given arabic number.
//If so, we calculate quotient by dividing our number by the number in compareSlice and also take the remainder of this operation.
//Quotient is used to repeat desired roman numeral for this position as many times as needed: 2000 >> MM.
//This roman numeral is then appended to resulting string.
//If remainder is not 0, recursiveBuilder is called again with the remainder as its parameter.
func (res *Roman) recursiveBuilder(arabic int) {
	remainder := 0
	for _, element := range compareSlice {
		if element <= arabic {
			quotient := arabic / element
			remainder = arabic % element
			*res += Roman(strings.Repeat(numbersToRoman[element], quotient))
			break
		}
	}
	switch remainder {
	case 0:
		break
	default:
		res.recursiveBuilder(remainder)
	}
}
