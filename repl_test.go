package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "bulbasaur CHARMANDER sqUirtle",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %q; want %q", c.input, actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%q) = %q; want %q", c.input, actual, c.expected)
			}
		}
	}
}

func TestCommandRegistry(t *testing.T) {
	cases := []struct {
		commandName string
		expected    string
	}{
		{
			commandName: "help",
			expected:    "Display a help message",
		},
		{
			commandName: "exit",
			expected:    "Exit the Pokedex",
		},
	}

	for _, c := range cases {
		cmd, exists := getCommands()[c.commandName]
		if !exists {
			t.Errorf("command %q does not exist in the commands map", c.commandName)
			continue
		}
		if cmd.description != c.expected {
			t.Errorf("command %q has description %q; want %q", c.commandName, cmd.description, c.expected)
		}
	}
}
