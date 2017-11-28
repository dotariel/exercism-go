package allergies

import "math"

var substances = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) (allergies []string) {
	for i, substance := range substances {
		if value := uint(math.Pow(2, float64(i))); (value & score) != 0 {
			allergies = append(allergies, substance)
		}
	}
	return
}

func AllergicTo(score uint, substance string) bool {
	if idx := indexOf(substance); idx > -1 {
		return uint(math.Pow(2, float64(idx)))&score != 0
	}
	return false
}

func indexOf(substance string) int {
	for i, s := range substances {
		if s == substance {
			return i
		}
	}
	return -1
}
