package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "My NaMe is JOhn",
			expected: []string{"my", "name", "is", "john"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of the slice not matching. Expected: %v, got: %v", len(c.expected), len(actual))
			continue
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words not matching. Expected: %v, got: %v", expectedWord, word)
				t.Fail()
			}
		}
	}
}
