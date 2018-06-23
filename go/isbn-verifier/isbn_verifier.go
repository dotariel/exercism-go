package isbn

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidISBN(s string) bool {
	isbn := strip(s)

	if len(isbn) != 10 {
		return false
	}

	var valid = regexp.MustCompile(`^[0-9]{9}[0-9X]$`)
	if !valid.MatchString(isbn) {
		return false
	}

	total := 0
	for i, v := range isbn {
		pos := int(i + 1)

		var mlt = ((len(isbn) - (pos)) % len(isbn)) + 1
		var val = 0

		if pos == len(isbn) && v == 'X' {
			val = 10
		} else {
			val = makeInt(v)
		}
		total += val * mlt
	}

	return total%11 == 0
}

func makeInt(r rune) int {
	t, _ := strconv.ParseInt(string(r), 10, 32)
	return int(t)
}

func strip(s string) string {
	return strings.Replace(s, "-", "", -1)
}
