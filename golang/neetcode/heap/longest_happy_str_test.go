package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestDiverseString(t *testing.T) {
	// Helper function to validate if a string is a happy string (no three consecutive same characters)
	isHappyString := func(s string) bool {
		if len(s) < 3 {
			return true
		}
		for i := 0; i <= len(s)-3; i++ {
			if s[i] == s[i+1] && s[i+1] == s[i+2] {
				return false
			}
		}
		return true
	}

	// Helper function to check if the result uses the correct number of characters
	checkCharCounts := func(result string, a, b, c int) bool {
		counts := map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
		}
		for _, ch := range result {
			counts[ch]++
		}
		// Check that we didn't use more than available
		if counts['a'] > a || counts['b'] > b || counts['c'] > c {
			return false
		}
		mx, n := max(a, b, c), a+b+c
		return len(result) == min(n, (n-mx)*3+2) // use as many mx as possible
	}

	tests := []struct {
		name        string
		a, b, c     int
		description string
	}{
		{"LeetCode Example 1", 1, 1, 7, "One a, one b, seven c's"},
		{"LeetCode Example 2", 7, 1, 0, "Seven a's, one b, zero c's"},
		{"All zeros", 0, 0, 0, "No characters"},
		{"Single character", 5, 0, 0, "Only a's (max two consecutive)"},
		{"Two characters equal", 3, 3, 0, "Three a's, three b's"},
		{"All characters equal", 3, 3, 3, "Three of each character"},
		{"Exact maximum allowed", 2, 2, 2, "Two of each character"},
		{"Dominant character", 4, 1, 1, "Four a's, one b, one c"},
		{"Just below limit", 3, 3, 3, "Three of each character"},
		{"Edge case with two zeros", 0, 0, 5, "Only c's (max two consecutive)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestDiverseString(tt.a, tt.b, tt.c)

			// Test if the result is a valid happy string
			assert.True(t, isHappyString(result), "Result should be a happy string: %s", result)

			// Test if the result uses characters correctly
			assert.True(t, checkCharCounts(result, tt.a, tt.b, tt.c), "Result doesn't use characters correctly: %s", result)

			// For special cases with known expected lengths
			switch tt.name {
			case "All zeros":
				assert.Empty(t, result, "Result should be empty when all counts are zero")
			case "Single character", "Edge case with two zeros":
				assert.Len(t, result, 2, "Result should have maximum 2 characters when only one type is available")
			}
		})
	}
}
