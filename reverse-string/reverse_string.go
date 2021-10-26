package reverse

//not an original solution by me, but I wanted to type it out and drill into my brain that I can easily swap slice lements around
func Reverse(input string) string {
	sliced := []rune(input)
	for i, j := 0, len(sliced)-1; i < j; i, j = i+1, j-1 {
		sliced[i], sliced[j] = sliced[j], sliced[i]
	}
	return string(sliced)
}
