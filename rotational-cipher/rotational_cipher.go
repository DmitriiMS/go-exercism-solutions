package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	newstr := make([]rune, len(plain))
	for i, c := range plain {
		if unicode.IsLetter(c) {
			if unicode.ToLower(c)-'a'+rune(shiftKey) >= 26 {
				newstr[i] = c + (rune(shiftKey) - 26)
				continue
			}
			newstr[i] = c + rune(shiftKey)
			continue
		}
		newstr[i] = c
	}
	return string(newstr)
}
