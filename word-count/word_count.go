package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	freq := make(Frequency)
	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && (c) != '\''
	}
	words := strings.FieldsFunc(input, splitter)
	for _, word := range words {
		if len(word) > 0 {
			freq[strings.ToLower(strings.Trim(string(word), "'"))] += 1
		}
	}
	return freq
}
