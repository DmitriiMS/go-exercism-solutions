// Package scrabble provides scrabble score calculation
package scrabble

import "strings"

// LetterScore returns a score of a single letter
func LetterScore(c string) int {
	switch c {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	default:
		return 0
	}
}

// Score calculates the scrabble score of a word
func Score(word string) (score int) {
	for _, char := range word {
		score += LetterScore(strings.ToUpper(string(char)))
	}
	return
}
