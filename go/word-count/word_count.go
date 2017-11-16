package wordcount

import (
	"strings"
	"unicode"
)

const (
	apostrophe rune = rune(39)
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var f = Frequency{}
	for _, word := range strings.FieldsFunc(phrase, split) {
		clean := strip(clean(word))
		if len(clean) > 0 {
			f[clean]++
		}
	}

	return f
}

func strip(word string) string {
	return strings.Trim(word, string(apostrophe))
}

func clean(word string) string {
	return strings.ToLower(strings.Map(sanitize, word))
}

func sanitize(r rune) rune {
	if unicode.In(r, unicode.Letter, unicode.Number) || r == apostrophe {
		return r
	}
	return -1
}

func split(r rune) bool {
	if !unicode.In(r, unicode.Space, unicode.Po) {
		return false
	}
	return r != apostrophe
}
