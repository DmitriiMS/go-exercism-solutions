//package anagram provides function to detect anagrams
package anagram

import (
	"sort"
	"strings"
)

//Detect detects anagrams
func Detect(word string, candidates []string) (result []string) {
	upWord := strings.ToUpper(word)               //uppercase the target word because anagram search should be case insensitive
	candidates = checkStrings(upWord, candidates) //candidate validation
	if len(candidates) == 0 {                     //if slice is empty after the validation, return empty slice
		return result
	}
	sortedWord := SortString(upWord) //sort the target word by characters
	//uppercase each candidate, sort it by letters and if it equals to the target word, add it to the resulting slice
	for _, candidate := range candidates {
		if SortString(strings.ToUpper(candidate)) != sortedWord {
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

//SortString sorts string by letters
func SortString(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}
