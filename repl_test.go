package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "", // empty string
			expected: []string{},
		},
		{
			input:    "	", // tab
			expected: []string{},
		},

		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Whacky snOO snoo!",
			expected: []string{"whacky", "snoo", "snoo!"},
		},
		{
			input:    "Beep, Boop,  Bop  ",
			expected: []string{"beep,", "boop,", "bop"},
		},

	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("len \"%s\" = %d; want len = %d", actual, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]

			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word %s; want %s", word, expectedWord)
			}

		}

	}
}
