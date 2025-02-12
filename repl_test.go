package main

import (
	"reflect"
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
		},
		{
			input:    "THIS IS A     TesT",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
	}
}
