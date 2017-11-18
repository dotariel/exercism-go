package anagram

import (
	"fmt"
	"sort"
	"testing"
)

func equal(a []string, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestDetectAnagrams(t *testing.T) {
	for _, tt := range testCases {
		actual := Detect(tt.subject, tt.candidates)
		if !equal(tt.expected, actual) {
			msg := `FAIL: %s
	Subject %s
	Candidates %q
	Expected %q
	Got %q
				`
			t.Fatalf(msg, tt.description, tt.subject, tt.candidates, tt.expected, actual)
		} else {
			t.Logf("PASS: %s", tt.description)
		}
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {

	b.StopTimer()

	for _, tt := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Detect(tt.subject, tt.candidates)
		}

		b.StopTimer()
	}

}

func TestAnagram(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{"tan", "ant", true},
		{"listen", "inlets", true},
		{"listen", "enlists", false},
		{"galea", "eagle", false},
		{"dog", "goody", false},
		{"allergy", "gallery", true},
		{"allergy", "ballerina", false},
		{"allergy", "regally", true},
		{"allergy", "clergy", false},
		{"allergy", "largely", true},
		{"allergy", "leading", false},
	}

	for _, tt := range cases {
		if actual := IsAnagram(tt.a, tt.b); tt.expected != actual {
			t.Errorf("'%v <-> %v' failed anagram test; wanted:%v got:%v", tt.a, tt.b, tt.expected, actual)
		}
	}
}
