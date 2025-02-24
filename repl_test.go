package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		expect []string
	}{
		{
			input:  "  hello  world   ",
			expect: []string{"hello", "world"},
		},
		{
			input:  "    ChaRMandER     BulBaSaur     sqUirtlE  pikaCHu    ",
			expect: []string{"charmander", "bulbasaur", "squirtle", "pikachu"},
		},
		{
			input:  "   tHEre is    bridge",
			expect: []string{"there", "is", "bridge"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expect) {
			t.Errorf("Length mismatch: expected %v, got %v", len(c.expect), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expect[i]
			if word != expectedWord {
				t.Errorf("Word mismatch at index %d: expected %q, got %q", i, expectedWord, word)
			}

		}
	}

}
