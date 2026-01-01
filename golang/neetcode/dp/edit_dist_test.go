package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		// LeetCode examples
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		
		// Edge cases
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"abc", "abc", 0},
		
		// Simple cases
		{"a", "b", 1},
		{"ab", "ac", 1},
		{"kitten", "sitting", 3},
		
		// Complex cases
		{"sunday", "saturday", 3},
		{"food", "money", 4},
	}

	for _, tt := range tests {
		t.Run(tt.a+"_"+tt.b, func(t *testing.T) {
			result := minDistance(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}