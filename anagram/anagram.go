//package anagram provides function to detect anagrams
package anagram

//following is not my original solution, it's an idea of "scoring" words in a way by using arrays
import (
	"strings"
)

//Detect detects anagrams
func Detect(word string, candidates []string) (result []string) {
	upWord := strings.ToUpper(word) //uppercase the target word
	wordScore := Score(upWord)      //calculate score for target word
	for _, candidate := range candidates {
		upCandidate := strings.ToUpper(candidate) //uppercase each candidate
		if upCandidate == upWord || wordScore != Score(upCandidate) {
			//if word and candidate are equal, or if their score are not equal, go to the next candidate
			continue
		}
		result = append(result, candidate) //append anagram to resulting slice
	}
	return
}

//Score generates an array of 26 elements, one for each letter,
//which serves as a mask for a give word, keeping the information about
//how many times each letter was used in this word.
//Those masks can be used as scores for words, anagrams should have same masks-scores
func Score(word string) (score [26]int) {
	for _, r := range word {
		if r-65 >= 0 { //check that we only iterate over letters, no spaces allowed
			score[r-65]++
		}
	}
	return
}
