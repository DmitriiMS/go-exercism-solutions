package etl

import (
	"strings"
)

type old map[int][]string
type new map[string]int

func Transform(in old) new {
	var transformed = make(new)
	for k, v := range in {
		for _, letter := range v {
			transformed[strings.ToLower(letter)] = k
		}
	}
	return transformed
}
