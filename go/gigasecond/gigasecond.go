package gigasecond

import (
	"time"
)

const testVersion = 4
const Gigasecond = "1000000000s"

func AddGigasecond(t time.Time) time.Time {
	duration, err := time.ParseDuration(Gigasecond)

	if err != nil {
		panic(err)
	}

	t = t.Add(duration)

	return t
}
