package twelve

import (
	"fmt"
)

const testVersion = 1

func Verse(day int) string {

	days := [][]string{
		[]string{"first", "a Partridge in a Pear tree"},
		[]string{"second", "two Turtle Doves"},
		[]string{"third", "three French Hens"},
		[]string{"fourth", "four Calling Birds"},
		[]string{"fifth", "five Gold Rings"},
		[]string{"sixth", "six Geese-a-Laying"},
		[]string{"seventh", "seven Swans-a-Swimming"},
		[]string{"eighth", "eight Maids-a-Milking"},
		[]string{"ninth", "nine Ladies Dancing"},
		[]string{"tenth", "ten Lords-a-Leaping"},
		[]string{"eleventh", "eleven Pipers Piping"},
		[]string{"twelfth", "twelve Drummers Drumming"},
	}

	index := day - 1

	verse := fmt.Sprintf("On the %v day of Christmas my true love gave to me, ", days[index][0])

	for i := index; i > 0; i-- {
		verse += fmt.Sprintf("%v, ", days[i][1])
	}
	if day > 1 {
		verse += "and "
	}
	verse += "a Partridge in a Pear Tree."

	return verse
}

func Song() string {
	var song string

	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}

	return song
}
