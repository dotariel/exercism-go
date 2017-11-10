package grains

import (
	"errors"
)

// Square returns the number of grains at the given position on the 8x8 board
func Square(in int) (uint64, error) {
	if in < 1 || in > 64 {
		return 0, errors.New("Input must be between 1 and 64")
	}

	return 1 << uint(in-1), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	return 1<<64 - 1
}
