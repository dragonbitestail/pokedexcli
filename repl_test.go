package main

import (
	"testing"
	"io"
	"strings"
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

// Once, refactored startRepl() to allow passing of io.Reader instead of hardcoding
// os.Stdin
//   (https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it),
// and handling special isTest condition along with exitState,
// we can test.
// !!!! NOTE: Any tests which call startRepl() MUST pass true as third param to tell
// repl loop to exit on receiving "exit" command or encountering command errors.
func TestStartRepl(t *testing.T){
	isTest := true

	//input := "foo\ncatch\ncatch 33\nexit\n"  // Produces error due to bad cacth w/o param.
	input := "foo\ncatch 33\nexit\n"
	var reader io.Reader = strings.NewReader(input)

	err, exitState := startRepl(&cfg, reader, isTest)
	if err != nil || ! exitState.exit {
		t.Errorf("startRepl() error: %v", err)
	}

	// Conversely, to check for a specific error, eval. the error message:
	want := "catch command requires Pokemon name parameter: catch <pokemon>"
	input2 := "catch\n"  // Produces error due to bad cacth w/o param.
	reader = strings.NewReader(input2)
	err, exitState = startRepl(&cfg, reader, isTest)
	if err != nil || ! exitState.exit {
		got := err.Error()
		if got != want {
			t.Errorf("startRepl() error: %v", err)
		}
	}

}
