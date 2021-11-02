package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	freq := make(Frequency)
	wordFinder := regexp.MustCompile(`[\w\d']+`)
	words := wordFinder.FindAll([]byte(input), -1)
	for _, word := range words {
		if len(word) > 0 {
			freq[strings.ToLower(strings.Trim(string(word), "'"))] += 1
		}
	}
	return freq
}
