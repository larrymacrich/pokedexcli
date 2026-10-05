package cli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"happy path simple": {
			input:    "charmander",
			expected: []string{"charmander"},
		},
		"happy path sequence": {
			input:    "charmander bulbasaur",
			expected: []string{"charmander", "bulbasaur"},
		},
		"sequence of spaces": {
			input:    "     ",
			expected: []string{},
		},
		"empty": {
			input:    "",
			expected: []string{},
		},
		"captial case": {
			input:    "Charmander",
			expected: []string{"charmander"},
		},
		"all caps": {
			input:    "CHARMANDER",
			expected: []string{"charmander"},
		},
		"mixed caps": {
			input:    "chARmaNDER",
			expected: []string{"charmander"},
		},
		"captial case sequence": {
			input:    "Charmander Bulbasaur",
			expected: []string{"charmander", "bulbasaur"},
		},
		"trailing spaces with capital case sequence": {
			input:    "  Charmander        Bulbasaur  ",
			expected: []string{"charmander", "bulbasaur"},
		},
		"mixed caps sequence": {
			input:    "CHARMANDER bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	// Loop over test cases
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Fatalf("expected: %v, got: %v", c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Fatalf("expected: %v, got: %v", expectedWord, word)
				continue
			}
		}
	}
}
