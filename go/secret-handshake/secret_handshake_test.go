package secret

import (
	"reflect"
	"testing"
)

const targetTestVersion = 2

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		input    uint
		expected []uint
	}{
		{1, []uint{1}},
		{2, []uint{0, 1}},
		{3, []uint{1, 1}},
		{19, []uint{1, 1, 0, 0, 1}},
	}

	for _, data := range cases {
		conversion := Convert(data.input)
		expected := data.expected
		if !reflect.DeepEqual(conversion, expected) {
			t.Fatalf("Input: %v, Expected: %v; Actual:%v", data.input, expected, conversion)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"wink", "foo"}, []string{"foo", "wink"}},
		{[]string{"", ""}, []string{"", ""}},
	}

	for _, data := range cases {
		reversed := Reverse(data.input)
		expected := data.expected
		if !reflect.DeepEqual(reversed, expected) {
			t.Fatalf("Input: %v, Expected: %v; Actual:%v", data.input, expected, reversed)
		}
	}
}

func TestHandshake(t *testing.T) {
	for _, test := range tests {
		h := Handshake(test.code)
		// use len() to allow either nil or empty list, because
		// they are not equal by DeepEqual
		if len(h) == 0 && len(test.h) == 0 {
			continue
		}
		if !reflect.DeepEqual(h, test.h) {
			t.Fatalf("Handshake(%d) = %q, want %q.", test.code, h, test.h)
		}
	}
}

func BenchmarkHandshake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Handshake(31)
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(19)
	}
}
