package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		expected int
	}{
		// Basic cases
		{
			name:     "Single car",
			target:   10,
			position: []int{0},
			speed:    []int{1},
			expected: 1,
		},
		{
			name:     "Two cars same speed - fleet forms",
			target:   10,
			position: []int{0, 5},
			speed:    []int{1, 1},
			expected: 2,
		},
		{
			name:     "Two cars different speeds - separate fleets",
			target:   10,
			position: []int{0, 6},
			speed:    []int{2, 1},
			expected: 2,
		},
		{
			name:     "Three cars - two fleets",
			target:   12,
			position: []int{10, 8, 0},
			speed:    []int{2, 4, 1},
			expected: 2,
		},

		// Edge cases
		{
			name:     "No cars",
			target:   10,
			position: []int{},
			speed:    []int{},
			expected: 0,
		},
		{
			name:     "Cars already at target",
			target:   10,
			position: []int{10, 10},
			speed:    []int{1, 2},
			expected: 1,
		},
		{
			name:     "Cars at same position different speeds",
			target:   10,
			position: []int{5, 5},
			speed:    []int{1, 2},
			expected: 2,
		},

		// Complex cases
		{
			name:     "Multiple cars",
			target:   12,
			position: []int{0, 4, 2},
			speed:    []int{2, 1, 1},
			expected: 2,
		},
		{
			name:     "Mixed speeds and positions",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			expected: 1,
		},
		{
			name:     "Real world example",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := carFleet(tt.target, tt.position, tt.speed)
			assert.Equal(t, tt.expected, result,
				"carFleet(target=%d, position=%v, speed=%v) = %d, expected %d",
				tt.target, tt.position, tt.speed, result, tt.expected)
		})
	}
}
