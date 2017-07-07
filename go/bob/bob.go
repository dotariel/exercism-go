package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(statement string) string {

	s := strings.TrimSpace(statement)

	if isYelling(s) {
		return "Whoa, chill out!"
	}

	if isQuestion(s) {
		return "Sure."
	}

	if isEmpty(s) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isYelling(s string) bool {
	hasLowercase, _ := regexp.MatchString("[a-z]+", s)
	hasUppercase, _ := regexp.MatchString("[A-Z]{2,}", s)

	return hasUppercase && !hasLowercase
}

func isEmpty(s string) bool {
	return s == ""
}

func contains(s []rune, c rune) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}
	return false
}
