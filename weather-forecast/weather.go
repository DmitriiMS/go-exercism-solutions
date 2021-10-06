// Package or something, idk, lol
package weather

var (
	// This is the current state of condition
	CurrentCondition string
	// Where am I, who am I, why am I?
	CurrentLocation string
)

// Log, like a round wooden log. Overall, not a great exercise, really.
func Log(city, condition string) string {

	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
