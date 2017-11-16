package phonenumber

import (
	"fmt"
	"regexp"
	"unicode"
)

func Number(in string) (string, error) {
	res := strip(in)
	rgx := "^[+]?[1]?[2-9][0-9]{2}[2-9][0-9]{6}$"

	if valid, _ := regexp.MatchString(rgx, res); !valid {
		return res, fmt.Errorf("invalid format")
	}

	if res[0:1] == "1" {
		res = res[1:len(res)]
	}

	return res, nil
}

func AreaCode(in string) (string, error) {
	number, err := Number(in)

	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

func Format(in string) (string, error) {
	number, err := Number(in)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%v) %v-%v", number[0:3], number[3:6], number[6:10]), nil
}

func strip(s string) string {
	var stripped string

	for _, char := range s {
		if unicode.IsDigit(char) {
			stripped += string(char)
		}
	}

	return stripped
}
