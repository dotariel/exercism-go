package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return -1, fmt.Errorf("cannot be zero")
	}
	return reduce(n, 0), nil
}

func reduce(n, count int) int {
	if n == 1 {
		return count
	}
	if n%2 == 0 {
		return reduce(n/2, count+1)
	}
	return reduce((n*3)+1, count+1)
}
