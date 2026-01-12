package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorganizeString(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		shouldBeValid bool
	}{
		{
			name:          "LeetCode Example 1",
			input:         "aab",
			shouldBeValid: true,
		},
		{
			name:          "LeetCode Example 2",
			input:         "aaab",
			shouldBeValid: false,
		},
		{
			name:          "Single character",
			input:         "a",
			shouldBeValid: true,
		},
		{
			name:          "Two different characters",
			input:         "ab",
			shouldBeValid: true,
		},
		{
			name:          "Two same characters",
			input:         "aa",
			shouldBeValid: false,
		},
		{
			name:          "Three characters with valid rearrangement",
			input:         "aabb",
			shouldBeValid: true,
		},
		{
			name:          "All characters same - impossible",
			input:         "aaaa",
			shouldBeValid: false,
		},
		{
			name:          "Multiple characters with valid solution",
			input:         "aabbbc",
			shouldBeValid: true,
		},
		{
			name:          "Complex case with many characters",
			input:         "aaabbcc",
			shouldBeValid: true,
		},
		{
			name:          "Characters with equal frequency",
			input:         "abcabc",
			shouldBeValid: true,
		},
		{
			name:          "Large frequency difference",
			input:         "aaaabbbcc",
			shouldBeValid: true,
		},
		{
			name:          "Impossible case with dominant character",
			input:         "aaaaabbb",
			shouldBeValid: false,
		},
		{
			name:          "Mixed case characters",
			input:         "aA",
			shouldBeValid: true,
		},
		{
			name:          "Special characters",
			input:         "!!??",
			shouldBeValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reorganizeString(tt.input)

			if tt.shouldBeValid {
				// Verify the result is not empty
				assert.NotEqual(t, "", result, "Result should not be empty for valid cases")
				// Verify the result has the same length as input
				assert.Equal(t, len(tt.input), len(result), "Result should have same length as input")
				// Verify no two adjacent characters are the same
				for i := 0; i < len(result)-1; i++ {
					assert.NotEqual(t, result[i], result[i+1], "Adjacent characters should not be the same")
				}
				// Verify all characters from input are present in result
				inputChars := make(map[rune]int)
				resultChars := make(map[rune]int)
				for _, ch := range tt.input {
					inputChars[ch]++
				}
				for _, ch := range result {
					resultChars[ch]++
				}
				assert.Equal(t, inputChars, resultChars, "Result should contain all characters from input with same frequency")
			} else {
				assert.Equal(t, "", result, "Result should be empty for impossible cases")
			}
		})
	}
}
