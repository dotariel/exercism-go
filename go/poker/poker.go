package poker

import (
	"errors"
)

const testVersion = 5

func BestHand(hands []string) ([]string, error) {

	var err error
	c := make([]string, 0, 0)

	if len(hands) == 0 {
		err = errors.New("No hands provided")
	}

	if len(hands) == 1 {
		c = append(c, hands[0])
	}

	return c, err
}
