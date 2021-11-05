package atbash

import "unicode"

type message struct {
	str string
	i   int
}

func Atbash(s string) string {
	encoded := message{"", 0}
	for _, c := range s {
		if unicode.IsLetter(c) {
			swapped := 'a' + 25 - (unicode.ToLower(c) - 'a')
			encoded.addChar(swapped)
		} else if unicode.IsDigit(c) {
			encoded.addChar(c)
		}
	}
	return encoded.str
}

func (m *message) addChar(char rune) {
	if m.i == 5 {
		m.str += " "
		m.i = 0
	}
	m.str += string(char)
	m.i++
}
