package pangram

import (
	"strings"
)

const testVersion = 1
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(s string) bool {
	found := 0

	for _, c := range alphabet {
		if inAlphabet(s, c) {
			found++
		}
	}

	return found == len(alphabet)
}

func inAlphabet(s string, c rune) bool {
	return strings.ContainsRune(s, c) || strings.Contains(s, strings.ToUpper(string(c)))
}
