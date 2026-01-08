package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakesquare(t *testing.T) {
	tests := []struct {
		name        string
		matchsticks []int
		expected    bool
	}{{
		name:        "Basic case - can form square",
		matchsticks: []int{1, 1, 2, 2, 2},
		expected:    true,
	}, {
		name:        "Cannot form square - sum not divisible by 4",
		matchsticks: []int{1, 1, 1, 1, 1},
		expected:    false,
	}, {
		name:        "Cannot form square - individual matchstick too long",
		matchsticks: []int{3, 3, 3, 3, 4},
		expected:    false,
	}, {
		name:        "Edge case - exactly 4 matchsticks of same length",
		matchsticks: []int{5, 5, 5, 5},
		expected:    true,
	}, {
		name:        "Edge case - only 3 matchsticks",
		matchsticks: []int{1, 2, 3},
		expected:    false,
	}, {
		name:        "Edge case - zero length matchstick",
		matchsticks: []int{0, 0, 0, 0},
		expected:    true,
	}, {
		name:        "Complex case - cannot form square",
		matchsticks: []int{2, 4, 4, 5, 6, 7, 8, 9, 17, 18},
		expected:    false,
	}, {
		name:        "Complex case - can form square",
		matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		expected:    true,
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := makesquare(tt.matchsticks)
			assert.Equal(t, tt.expected, result, "makesquare(%v) should be %v", tt.matchsticks, tt.expected)
		})
	}
}
