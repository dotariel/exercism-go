package etl

import (
	"strings"
)

// Transform converts the legacy format of the Scrabble scoring system
// into the new format.
func Transform(m map[int][]string) map[string]int {
	res := make(map[string]int)
	for i, ss := range m {
		for _, letter := range ss {
			res[strings.ToLower(letter)] = i
		}
	}
	return res
}
