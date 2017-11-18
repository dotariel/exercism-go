package anagram

import (
	"sort"
	"strings"
)

type byRunes []rune

func (s byRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s byRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byRunes) Len() int {
	return len(s)
}

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(byRunes(runes))
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)

	for _, candidate := range candidates {
		if IsAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func IsAnagram(a, b string) bool {
	this := strings.ToLower(a)
	that := strings.ToLower(b)

	if this == that {
		return false
	}

	return sorted(this) == sorted(that)
}
