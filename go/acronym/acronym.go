package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Abbreviate(s string) string {
	acronym := ""
	for _, v := range strings.FieldsFunc(s, split) {
		acronym += strings.ToUpper(string(v[0]))
	}

	return acronym
}

func split(c rune) bool {
	return unicode.In(c, unicode.Dash, unicode.Space)
}
