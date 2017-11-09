package luhn

import (
	"strconv"
	"unicode"
)

func Valid(s string) bool {
	digits := make([]int64, 0)

	for _, d := range s {
		if !(unicode.IsDigit(d) || unicode.IsSpace(d)) {
			return false
		}

		if !unicode.IsSpace(d) {
			digit, _ := strconv.ParseInt(string(d), 10, 64)
			digits = append(digits, digit)
		}
	}

	if len(digits) <= 1 {
		return false
	}

	var total int64
	double := false
	for i := len(digits) - 1; i >= 0; i-- {
		if double {
			t := digits[i] * 2
			if t > 9 {
				t = t - 9
			}
			total += t
			double = false
		} else {
			total += digits[i]
			double = true
		}
	}

	return total%10 == 0
}
