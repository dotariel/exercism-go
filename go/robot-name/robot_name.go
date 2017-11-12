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

const (
	charOffset  = 65
	charLength  = 26
	digitOffset = 49
	digitLength = 10
)

func makename() string {
	return gen(2, (charLength-1), charOffset) + gen(3, (digitLength-1), digitOffset)
}

func gen(count, len, offset int) string {
	var name string
	for i := 0; i < count; i++ {
		name += string(rune(rand.Intn(len) + offset))
	}
	return name
}
