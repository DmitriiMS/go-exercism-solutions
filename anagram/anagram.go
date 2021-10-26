//package anagram provides function to detect anagrams
package anagram

import (
	"sort"
	"strings"
)

type ByRune []rune //custom type for methods

//methods needed to use sort function.
func (r *ByRune) Len() int           { return len(*r) }
func (r *ByRune) Swap(i, j int)      { (*r)[i], (*r)[j] = (*r)[j], (*r)[i] }
func (r *ByRune) Less(i, j int) bool { return (*r)[i] < (*r)[j] }

//Detect detects anagrams
func Detect(word string, candidates []string) (result []string) {
	upWord := strings.ToUpper(word)               //uppercase the target word because anagram search should be case insensitive
	candidates = checkStrings(upWord, candidates) //candidate validation
	if len(candidates) == 0 {                     //if slice is empty after the validation, return empty slice
		return result
	}
	sortedWord := SortStringByCharacter(upWord) //sort the target word by characters
	//uppercase each candidate, sort it by letters and if it equals to the target word, add it to the resulting slice
	for _, candidate := range candidates {
		if SortStringByCharacter(strings.ToUpper(candidate)) != sortedWord {
			continue
		}
		result = append(result, candidate)
	}
	return result
}

//checkStrings checks if candidates are same as the target word
func checkStrings(word string, candidates []string) []string {
	//compare each uppercased candidate to target starting from the end of the slice
	//if the match, trim slice
	for i := len(candidates) - 1; i >= 0; i-- {
		if strings.ToUpper(candidates[i]) == word {
			candidates = candidates[:len(candidates)-1]
		}
	}
	return candidates
}

//SortStringByCharacter converts string to rune slice, sorts it and returns it as a string
func SortStringByCharacter(s string) string {
	var r ByRune = []rune(s)
	sort.Sort(&r)
	return string(r)
}
