package reverse

func Reverse(input string) (output string) {
	for _, letter := range input {
		output = string(letter) + output
	}
	return
}
