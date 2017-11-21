package lsproduct

import (
	"fmt"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(series string, span int) (int64, error) {
	var max int64

	if span > len(series) {
		return 0, fmt.Errorf("oops")
	}

	if span == 0 {
		return 1, nil
	}

	if span < 0 {
		return -1, fmt.Errorf("span must be positive")
	}

	for i := 0; i < len(series); i++ {
		start := i
		end := i + span

		if end > len(series) {
			break
		}

		newMax, err := calculate(series[start:end])
		if err != nil {
			return -1, err
		}

		if newMax > max {
			max = newMax
		}
	}

	return max, nil
}

func calculate(span string) (int64, error) {
	var product int64 = 1

	for _, d := range span {
		if !unicode.IsNumber(d) {
			return -1, fmt.Errorf("not a digit")
		}

		n, _ := strconv.ParseInt(string(d), 10, 64)
		product *= n
	}

	return product, nil
}
