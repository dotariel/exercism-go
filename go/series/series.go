package series

const testVersion = 2

func All(n int, s string) []string {
	res := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if i+n <= len(s) {
			res = append(res, string(s[i:i+n]))
		}
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	all := All(n, s)
	ok = len(all) > 0

	if ok {
		first = all[0]
	}

	return
}
