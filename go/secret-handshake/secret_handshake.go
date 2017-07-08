package secret

const testVersion = 2

func Handshake(n uint) []string {
	lex := []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
	}

	bs := make([]string, 0)
	bits := Convert(n)

	for bitPos, bit := range bits {
		if bit == 1 && bitPos < len(lex) {
			bs = append(bs, lex[bitPos])
		}
	}

	if len(bits) > 4 && bits[4] == 1 {
		bs = Reverse(bs)
	}

	return bs
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func Convert(n uint) []uint {
	a, b := n, n
	xs := make([]uint, 0)

	for {
		a, b = a/2, a%2
		xs = append(xs, b)

		if a == 0 {
			break
		}
	}

	return xs
}
