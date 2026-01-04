package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		expected  []int
	}{
		{
			name:      "LeetCode Example 1",
			asteroids: []int{5, 10, -5},
			expected:  []int{5, 10},
		},
		{
			name:      "LeetCode Example 2",
			asteroids: []int{8, -8},
			expected:  []int{},
		},
		{
			name:      "LeetCode Example 3",
			asteroids: []int{10, 2, -5},
			expected:  []int{10},
		},
		{
			name:      "All positive asteroids",
			asteroids: []int{1, 2, 3, 4},
			expected:  []int{1, 2, 3, 4},
		},
		{
			name:      "All negative asteroids",
			asteroids: []int{-1, -2, -3, -4},
			expected:  []int{-1, -2, -3, -4},
		},
		{
			name:      "Mixed with no collisions",
			asteroids: []int{-1, -2, 3, 4},
			expected:  []int{-1, -2, 3, 4},
		},
		{
			name:      "Complex collision chain",
			asteroids: []int{5, 10, -5, -10},
			expected:  []int{5},
		},
		{
			name:      "Smaller asteroid destroyed",
			asteroids: []int{10, -5},
			expected:  []int{10},
		},
		{
			name:      "Equal size collision",
			asteroids: []int{5, -5},
			expected:  []int{},
		},
		{
			name:      "Empty array",
			asteroids: []int{},
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asteroidCollision(tt.asteroids)
			assert.Equal(t, tt.expected, result)
		})
	}
}
