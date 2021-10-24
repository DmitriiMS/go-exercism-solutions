// Package leap provides a function that calculates if the year is a leap year
package leap

// IsLeapYear calculates if given year is a leap year
func IsLeapYear(year int) bool {
	if (year%100 == 0 && year%400 != 0) || year%4 != 0 {
		return false
	}
	return true
}
