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
			if _, ok := transformed[strings.ToLower(letter)]; !ok {
				transformed[strings.ToLower(letter)] = k
			}
		}
	}
	return transformed
}
