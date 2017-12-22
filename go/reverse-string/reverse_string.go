package reverse

func String(s string) (ret string) {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	for _, rune := range runes {
		ret += string(rune)
	}

	return ret
}
