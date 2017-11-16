package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		multiples := make(map[int]bool)
		for _, d := range divisors {
			if i%d == 0 {
				if _, added := multiples[i]; !added {
					sum += i
					multiples[i] = true
				}
			}
		}
	}

	return sum
}
