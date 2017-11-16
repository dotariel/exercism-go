package wordcount

import (
	"reflect"
	"testing"
)

// wordcount API
//
// func WordCount(phrase string) Frequency  // Implement this function.
// type Frequency map[string]int            // Using this return type.

func TestWordCount(t *testing.T) {
	for _, tt := range testCases {
		expected := tt.output
		actual := WordCount(tt.input)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("%s\n\tExpected: %v\n\tGot: %v", tt.description, expected, actual)
		} else {
			t.Logf("PASS: %s - WordCount(%s)", tt.description, tt.input)
		}
	}
}

func BenchmarkWordCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			WordCount(tt.input)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		input    rune
		expected bool
	}{
		{' ', true},
		{'.', true},
		{':', true},
		{apostrophe, false},
	}

	for _, tt := range cases {
		actual := split(tt.input)

		if actual != tt.expected {
			t.Errorf("scenario failed; input:[%v], expected:%v; got:%v", string(tt.input), tt.expected, actual)
		}
	}
}

func TestClean(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"don't", "don't"},
		{"javascript!!&@$%^&", "javascript"},
		{"1", "1"},
	}

	for _, tt := range cases {
		actual := clean(tt.input)

		if actual != tt.expected {
			t.Errorf("scenario failed; expected:%v; got:%v", tt.expected, actual)
		}
	}
}
