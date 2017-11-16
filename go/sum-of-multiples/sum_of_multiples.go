package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		if isMultipleOfAny(i, divisors...) {
			sum += i
		}
	}

	return
}

func isMultipleOfAny(n int, divisors ...int) bool {
	for _, d := range divisors {
		if n%d == 0 {
			return true
		}
	}

	return false
}
