package main

import (
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:"   palabra  2espacios   3espaCIOS    4espacios   ",
			expected: []string{"palabra","2espacios","3espacios","4espacios"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slices don't match")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words in slices don't match'")
				return
			}
		}
	}
}
