package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	runeSlice := []rune(word)
	for i, r1 := range runeSlice {
		if unicode.IsSpace(r1) || unicode.IsPunct(r1) {
			continue
		}
		for _, r2 := range runeSlice[i+1:] {
			if unicode.ToLower(r1) == unicode.ToLower(r2) {
				return false
			}
		}
	}
	return true
}
