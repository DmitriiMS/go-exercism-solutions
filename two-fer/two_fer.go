// Package twofer provideas a function that reurns a string with given name.
package twofer

import "fmt"

// ShareWith take a name and returns a string with it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
