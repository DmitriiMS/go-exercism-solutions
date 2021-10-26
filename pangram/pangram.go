//package pangram provides function to check if given sentence is a pangram
package pangram

import (
	"unicode"
)

//IsPangram takes a sentence and validates whether or not it is a pangram
func IsPangram(in string) bool {
	var codes []int = make([]int, 26) // create an int slice for all letters
	for i := 0; i <= 25; i++ {
		codes[i] = 1 //fill with ones, 1 means letter was not encountered, 0 -- was encountered
	}
	for _, letter := range in {
		uppercase := unicode.ToUpper(letter)         //take each rune in sentence and try to set it to upper case
		if uppercase-65 >= 0 && uppercase-65 <= 25 { //runes are just codes, in case of uppercase ascii letters -- codes from 65 to 90; substracting 65 gives a range of 0 to 25
			codes[uppercase-65] = 0 // if in range, set this code to 0 -- we met this particular letter
		}
	}
	//now run over the slice and check that every item is 0, if not, given sentence lacked some letters, in which case return false
	for i := 0; i <= 25; i++ {
		if codes[i] != 0 {
			return false
		}
	}
	return true
}
