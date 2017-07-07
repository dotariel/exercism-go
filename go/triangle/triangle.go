package triangle

import (
	"math"
)

const testVersion = 3

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case !validSides(a, b, c):
		fallthrough
	case a+b < c || a+c < b || c+b < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	}
	return Sca
}

func validSides(sides ...float64) bool {
	for _, s := range sides {
		if s <= 0 || math.IsInf(s, 1) || math.IsNaN(s) {
			return false
		}
	}
	return true
}
