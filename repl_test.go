package main

import (
	"testing"
	"io"
	"strings"
)

//var cacheDuration = time.Second * time.Duration(200)
//var cachePokeAPI = pokecache.NewCache(time.Duration(cacheDuration))


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

// TODO: Refactor startRepl() to allow passing of
// io.Reader instead of hardcoding os.Stdin
// https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
func TestStartRepl(t *testing.T){

	input := "catch 33\n"
	var reader io.Reader = strings.NewReader(input)
	isTest := true
	startRepl(&cfg, reader, isTest)
	//if err != nil {
	//   t.Errorf("Failed to read from strings.NewReader: %w", err)
	//}
	return
}
