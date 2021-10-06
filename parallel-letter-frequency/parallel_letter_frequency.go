package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// CFrequency walks over the text string it gets and sends resulting map back via the cnannel
func CFrequency(s string, out chan<- FreqMap, wg *sync.WaitGroup) {
	defer wg.Done()     // reduce waitgroup after all work is done
	out <- Frequency(s) //send map back
}

func ConcurrentFrequency(s []string) FreqMap {
	res := FreqMap{}                     //slice to store our maps
	frequencies := make(chan FreqMap, 8) //channel to get our maps back
	var wg sync.WaitGroup                // waitgroup so we can close the channel
	// Loop over the slice, increase waitgroup, launch goroutine
	wg.Add(len(s))
	for _, text := range s {
		go CFrequency(text, frequencies, &wg)
	}

	//closer -- waits for all workers to finish and closes channel, so the loop can break
	go func() {
		wg.Wait()
		close(frequencies)
	}()

	//wait for inputs from channel, walk through results and merge our maps into one, break when channel is closed
	for rm := range frequencies {
		for k, v := range rm {
			res[k] += v
		}
	}
	return res
}
