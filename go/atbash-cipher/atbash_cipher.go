package atbash

import (
	"strings"
	"unicode"
)

const (
	key string = "abcdefghijklmnopqrstuvwxyz"
)

func Atbash(plain string) string {
	return format(strings.Map(mapf, strip(plain)))
}

func mapf(r rune) rune {
	if !unicode.IsLower(r) {
		return r
	}

	cidx := (len(key) - strings.IndexRune(key, r)) - 1

	return rune(key[cidx])
}

func strip(s string) (res string) {
	for _, char := range strings.ToLower(s) {
		if unicode.In(char, unicode.Lower, unicode.Digit) {
			res += string(char)
		}
	}
	return
}

func format(cipher string) (res string) {
	for i, char := range cipher {
		res += string(char)
		if (i+1)%5 == 0 && i != len(cipher)-1 {
			res += string(32)
		}
	}
	return
}
