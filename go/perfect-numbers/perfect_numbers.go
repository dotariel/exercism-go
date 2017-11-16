package perfect

import (
	"errors"
	"math"
)

var ErrOnlyPositive = errors.New("should be positive")

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	aliquot := aliquot(n)

	if aliquot == n {
		return ClassificationPerfect, nil
	}

	if aliquot < n {
		return ClassificationDeficient, nil
	}

	if aliquot > n {
		return ClassificationAbundant, nil
	}

	return 0, nil
}

func factors(n int64) (factors []int64) {
	fs := make(map[int64]bool)

	for i := int64(1); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			r := n / i
			fs[r] = true
			fs[i] = true
		}
	}

	for k := range fs {
		factors = append(factors, k)
	}

	return
}

func aliquot(n int64) (sum int64) {
	for _, f := range factors(n) {
		if f != n {
			sum += f
		}
	}
	return
}
