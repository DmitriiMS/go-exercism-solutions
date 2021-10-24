// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

type Kind string

const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if NaNiFCheck(a, b, c) {
		return NaT
	}
	if a <= (b+c) && b <= (a+c) && c <= (a+b) {
		if a == b && b == c {
			return Equ
		} else if a == b || b == c || a == c {
			return Iso
		} else {
			return Sca
		}
	}
	return NaT
}

func NaNiFCheck(a, b, c float64) bool {
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	if (math.IsNaN(a) || a == pinf || a == ninf || a == 0) ||
		(math.IsNaN(b) || b == pinf || b == ninf || b == 0) ||
		(math.IsNaN(c) || c == pinf || c == ninf || c == 0) {
		return true
	}
	return false
}
