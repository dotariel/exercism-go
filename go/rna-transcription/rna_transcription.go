package strand

import "strings"

var nucleotideMap = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	return strings.Map(mapf, dna)
}

func mapf(r rune) rune {
	return nucleotideMap[r]
}
