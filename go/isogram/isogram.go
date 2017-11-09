package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	var seen = make(map[rune]bool)

	word := strings.ToLower(s)
	for _, c := range word {
		if _, ok := seen[c]; ok {
			return false
		}

		if unicode.IsLetter(c) {
			seen[c] = true
		}
	}
	return true
}
