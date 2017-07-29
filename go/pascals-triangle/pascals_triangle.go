package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	t := make([][]int, 0, n)
	t = append(t, []int{1})

	for i := 1; i < n; i++ {
		pr := t[i-1]             // previous row
		nl := len(pr) + 1        // length of new row
		nr := make([]int, 0, nl) // new row
		nr = append(nr, 1)
		for pos := 1; pos < nl; pos++ {
			val := pr[pos-1]
			if pos < len(pr) {
				val += pr[pos]
			}
			nr = append(nr, val)
		}
		t = append(t, nr)
	}
	return t
}
