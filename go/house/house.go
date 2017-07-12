package house

import (
	"fmt"
)

const testVersion = 1

type VersePart struct {
	Verb string
	Noun string
}

var parts []VersePart = []VersePart{
	VersePart{Noun: "the horse and the hound and the horn", Verb: "belonged to"},
	VersePart{Noun: "the farmer sowing his corn", Verb: "kept"},
	VersePart{Noun: "the rooster that crowed in the morn", Verb: "woke"},
	VersePart{Noun: "the priest all shaven and shorn", Verb: "married"},
	VersePart{Noun: "the man all tattered and torn", Verb: "kissed"},
	VersePart{Noun: "the maiden all forlorn", Verb: "milked"},
	VersePart{Noun: "the cow with the crumpled horn", Verb: "tossed"},
	VersePart{Noun: "the dog", Verb: "worried"},
	VersePart{Noun: "the cat", Verb: "killed"},
	VersePart{Noun: "the rat", Verb: "ate"},
	VersePart{Noun: "the malt", Verb: "lay in"},
	VersePart{Noun: "the house that Jack built"},
}

func Song() string {
	var verses string

	for i := 1; i <= len(parts); i++ {
		verses += Verse(i)

		if i < len(parts) {
			verses += "\n\n"
		}
	}

	return verses
}

func Verse(n int) string {
	s := ""
	return buildVerse(n, &s)
}

func buildVerse(n int, s *string) string {
	if n == 0 {
		return fmt.Sprintf("This is %v.", *s)
	}

	part := parts[len(parts)-n]

	*s += part.Noun

	if part.Verb != "" {
		*s += fmt.Sprintf("\nthat %v ", part.Verb)
	}

	return buildVerse(n-1, s)
}
