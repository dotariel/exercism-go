package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(input string, shift int) string {
	var cipher string

	for _, char := range input {
		cipher += wrap(char, shift)
	}

	return cipher
}

func wrap(r rune, shift int) string {
	var base int

	if !unicode.IsLetter(r) {
		return string(r)
	}

	if unicode.IsLower(r) {
		base = int(byte('a'))
	}

	if unicode.IsUpper(r) {
		base = int(byte('A'))
	}

	return string(((int(r) + shift - base) % 26) + base)
}
