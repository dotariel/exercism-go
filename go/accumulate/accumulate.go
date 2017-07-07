package accumulate

const testVersion = 1

func Accumulate(s []string, op func(string) string) []string {
	result := make([]string, len(s))

	for i, val := range s {
		result[i] = op(val)
	}

	return result
}
