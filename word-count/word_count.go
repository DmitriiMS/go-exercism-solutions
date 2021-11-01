package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	freq := make(Frequency)
	punctuation := regexp.MustCompile(`[!?.,_\-:;&@$%^&\n\t]`)
	depunctuated := punctuation.ReplaceAll([]byte(input), []byte(" "))
	sliced := strings.Split(string(depunctuated), " ")
	for _, word := range sliced {
		if len(word) > 0 {
			freq[strings.ToLower(strings.Trim(word, "'"))] += 1
		}
	}
	return freq
}
