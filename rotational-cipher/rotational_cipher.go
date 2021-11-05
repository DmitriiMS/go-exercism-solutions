package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	newstr := make([]rune, len(plain))
	for i, c := range plain {
		switch {
		case c >= 'a' && c <= 'z':
			newstr[i] = 'a' + ((c - 'a' + rune(shiftKey)) % 26)
		case c >= 'A' && c <= 'Z':
			newstr[i] = 'A' + ((c - 'A' + rune(shiftKey)) % 26)
		default:
			newstr[i] = c
		}
	}
	return string(newstr)
}
