//package pangram provides function to check if given sentence is a pangram
package pangram

import (
	"unicode"
)

//IsPangram takes a sentence and validates whether or not it is a pangram
func IsPangram(in string) bool {
	var codes []int = make([]int, 91) // create an int slice with indecies from 0 to 90
	for i := 65; i < 91; i++ {        //uppercase letters from A to Z have ascii codes from 65 to 90
		codes[i] = 1 //fill indecies in this range with ones to represent that those letters are not yet encountered
	}
	for _, letter := range in {
		uppercase := unicode.ToUpper(letter)    //take each rune in sentence try to set it to upper case
		if uppercase >= 65 && uppercase <= 90 { //runes are just codes, in case of ascii -- ascii codes, here is a chek that it's a letter
			codes[uppercase] = 0 // if it's a letter, set this code to 0: already encountered
		}
	}
	//now run over the slice and check that every item is 0, if not, given sentence lacked some letters, in which case return false
	for i := 0; i < 91; i++ {
		if codes[i] != 0 {
			return false
		}
	}
	return true
}
