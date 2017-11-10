package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains at the given position in the 8x8 board
func Square(in int) (uint64, error) {
	if in < 1 || in > 64 {
		return 0, errors.New("Input must be between 1 and 64")
	}

	squares := make([]uint64, 64)
	pos := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			squares[pos] = uint64(math.Pow(float64(2), float64(pos)))

			// Return early if we have reached the desired position
			if pos == in-1 {
				return squares[pos], nil
			}

			pos++
		}
	}

	return squares[in-1], nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	var total uint64

	for i := 1; i < 65; i++ {
		t, _ := Square(i)
		total += t
	}

	return total
}
