package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP3804_SAM(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Simple duplicate substring",
			input:    "ababa",
			expected: 6, // "aba" appears 3 times: 3*3=9
		},
		{
			name:     "Single character repeated",
			input:    "aaaaa",
			expected: 9, // aaa * 3
		},
		{
			name:     "No duplicate substrings",
			input:    "abcd",
			expected: 0,
		},
		{
			name:     "Repeated pattern",
			input:    "ababab",
			expected: 8, // abab * 2
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 0, // Only one occurrence
		},
		{
			name:     "Two characters repeated",
			input:    "abab",
			expected: 4, // ab * 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input buffer
			input := bytes.NewBufferString(tt.input + "\n")
			// Create output buffer
			var output bytes.Buffer

			// Run the function
			P3804_SAM(input, &output)

			// Parse the output
			var result int
			_, err := fmt.Fscanf(&output, "%d", &result)
			assert.NoError(t, err)

			// Check the result
			assert.Equal(t, tt.expected, result)
		})
	}
}
