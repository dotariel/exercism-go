package perfect

import (
	"testing"
)

var _ error = ErrOnlyPositive

func TestGivesPositiveRequiredError(t *testing.T) {
	if _, err := Classify(0); err != ErrOnlyPositive {
		t.Fatalf("FAIL GivesPositiveRequiredError Expected error %q but got %q", ErrOnlyPositive, err)
	}
	t.Logf("PASS GivesPositiveRequiredError")
}

func TestClassifiesCorrectly(t *testing.T) {
	for _, c := range classificationTestCases {
		cat, err := Classify(c.input)
		switch {
		case err != nil:
			if c.ok {
				t.Fatalf("FAIL %s\nClassify(%d)\nExpected no error but got error %q", c.description, c.input, err)
			}
		case !c.ok:
			t.Fatalf("FAIL %s\nClassify(%d)\nExpected error but got %q", c.description, c.input, cat)
		case cat != c.expected:
			t.Fatalf("FAIL %s\nClassify(%d)\nExpected %q, got %q", c.description, c.input, c.expected, cat)
		}
		t.Logf("PASS %s", c.description)
	}
}

// Test that the classifications are not equal to each other.
// If they are equal, then the tests will return false positives.
func TestClassificationsNotEqual(t *testing.T) {
	classifications := []struct {
		class Classification
		name  string
	}{
		{ClassificationAbundant, "ClassificationAbundant"},
		{ClassificationDeficient, "ClassificationDeficient"},
		{ClassificationPerfect, "ClassificationPerfect"},
	}

	for i, pair1 := range classifications {
		for j := i + 1; j < len(classifications); j++ {
			pair2 := classifications[j]
			if pair1.class == pair2.class {
				t.Fatalf("%s should not be equal to %s", pair1.name, pair2.name)
			}
		}
	}
}

func BenchmarkClassify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range classificationTestCases {
			Classify(c.input)
		}
	}
}

func TestAliquot(t *testing.T) {
	cases := []struct {
		input    int64
		expected int64
	}{
		{1, 0},   // 1
		{2, 1},   // 1
		{3, 1},   // 1
		{4, 3},   // 1+2
		{12, 16}, // 1+2+3+4+6
		{15, 9},  // 1+3+5
		{18, 21}, // 1+2+3+6+9
		{24, 36}, // 1+2+3+4+6+8+12
		{28, 28}, // 1+2+4+7+14
	}

	for _, tt := range cases {
		actual := aliquot(tt.input)
		if actual != tt.expected {
			t.Errorf("failed for input '%v'; expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}
