package space

// Planet is a string alias
type Planet string

// Earth represents the baseline orbital age for Earth
const Earth float64 = 31557600

var conversions = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132,
}

// Age computes the age in a given Planet
func Age(a float64, p Planet) float64 {
	return round(a / conversions[p])
}

func round(x float64) float64 {
	return float64(int64(x/0.01+0.5)) * 0.01
}
