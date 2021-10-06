// Package hamming provides a function for calculating Hamming Distance between two strings
package hamming

import "errors"

// Distance function calculates the Hamming Distance between two strings
func Distance(a, b string) (int, error) {
	// checks first
	if len(a) != len(b) {
		return 0, errors.New("input strings have different lenghts")
	}
	if len(a) == 0 {
		return 0, nil
	}
	//assume that for this task ASCII is enough, can be genralised to use UTF runes
	counter := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
