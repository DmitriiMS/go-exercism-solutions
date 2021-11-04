package wordcount

import (
	"strings"

	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	frequency := make(Frequency)

	words := strings.FieldsFunc(phrase, func(r rune) bool {

		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')

	})

	for _, word := range words {

		word = strings.Trim(strings.ToLower(word), "'")

		if _, ok := frequency[word]; !ok {

			frequency[word] = 1

		} else {

			frequency[word]++

		}

	}

	return frequency

}
