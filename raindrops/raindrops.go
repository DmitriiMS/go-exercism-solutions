// Package raindrops provides plink-plank calculations
package raindrops

import "fmt"

// Converts plinks, planks and plongs
func Convert(num int) string {
	pling := ""
	plang := ""
	plong := ""
	if num%3 == 0 {
		pling = "Pling"
	}
	if num%5 == 0 {
		plang = "Plang"
	}
	if num%7 == 0 {
		plong = "Plong"
	}
	raindrops := pling + plang + plong
	if raindrops == "" {
		return fmt.Sprint(num)
	}
	return raindrops
}
