package raindrops

import (
	"strconv"
)

const testVersion = 3

func Convert(number int) string {
	conversion := ""

	if number%3 == 0 {
		conversion += "Pling"
	}

	if number%5 == 0 {
		conversion += "Plang"
	}

	if number%7 == 0 {
		conversion += "Plong"
	}

	if conversion == "" {
		conversion = strconv.Itoa(number)
	}

	return conversion
}
