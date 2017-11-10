package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	ch := make(chan FreqMap)

	for _, seq := range s {
		go func(s string) {
			ch <- Frequency(s)
		}(seq)
	}

	res := FreqMap{}

	for i := 0; i < len(s); i++ {
		m := <-ch
		for r, i := range m {
			res[r] += i
		}
	}

	return res
}
