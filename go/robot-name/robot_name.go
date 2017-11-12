package robotname

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var robots = make([]string, 0)

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if len(r.name) == 0 {
		for {
			n := makename()

			if e, _ := exists(n); !e {
				r.name = n
				robots = append(robots, n)
				break
			}
		}
	}

	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}

func exists(r string) (bool, int) {
	for i, robot := range robots {
		if r == robot {
			return true, i
		}
	}

	return false, -1
}

func makename() string {
	name := ""
	for i := 0; i < 2; i++ {
		name += string(rune(rand.Intn(25) + 65)) // Uppsercase 65-90
	}
	for i := 0; i < 3; i++ {
		name += string(rune(rand.Intn(9) + 49)) // Digits 49-58
	}

	return name
}
