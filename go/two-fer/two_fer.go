package twofer

import "fmt"

func ShareWith(s string) string {
	name := s

	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
