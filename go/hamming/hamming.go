package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {

	// Return right away if the strings are not comparable
	if len(a) != len(b) {
		return -1, errors.New("Strings must be of equal length")
	}

	// Loop through one string and compare the value at each
	// position with the other string, counting the differences.
	hamming := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil
}
