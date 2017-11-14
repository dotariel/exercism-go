package dna

import (
	"fmt"
	"strings"
)

const (
	nucleotides string = "ACGT"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Valid checks if a given rune represents a valid nucleotide
func Valid(n rune) bool {
	return strings.ContainsRune(nucleotides, n)
}

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	if !Valid(rune(nucleotide)) {
		return 0, fmt.Errorf("invalid nucleotide: %v", nucleotide)
	}

	count := 0
	for _, r := range d {
		if byte(r) == nucleotide {
			count++
		}
	}

	return count, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := make(map[byte]int)
	for _, n := range nucleotides {
		h[byte(n)] = 0
	}

	for _, r := range d {
		if !Valid(rune(r)) {
			return nil, fmt.Errorf("invalid nucleotide: %v", r)
		}
		h[byte(r)]++
	}

	return h, nil
}
