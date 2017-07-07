package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

type Clock struct {
	Hour, Minute int
}

func New(hour, minute int) Clock {
	return Clock{0, 0}.Add((hour * 60) + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {

	for i := 0; i < int(math.Abs(float64(minutes))); i++ {
		if minutes > 0 {
			c.Minute++

			if c.Minute == 60 {
				c.Hour++
				c.Minute = 0
			}
		}

		if minutes < 0 {
			if c.Hour == 0 {
				c.Hour = 24
			}

			if c.Minute == 0 {
				c.Minute = 60
				c.Hour--
			}

			c.Minute--
		}
	}

	c.Hour = c.Hour % 24

	return c
}
