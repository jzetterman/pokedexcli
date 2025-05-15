package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "Hello from earth",
			expected: []string{"hello", "from", "earth"},
		}, {
			input:    "Elon Musk is pretty cool",
			expected: []string{"elon", "musk", "is", "pretty", "cool"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d words",
				c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) word %d: got %q, expected %q",
					c.input, i, actual[i], c.expected[i])
			}
			fmt.Printf("cleanInput(%q) word %d: got %q, expected %q\n",
				c.input, i, actual[i], c.expected[i])
		}
	}
}
