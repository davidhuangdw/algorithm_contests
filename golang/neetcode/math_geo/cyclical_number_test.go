package math_geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "happy number 19",
			n:        19,
			expected: true,
		},
		{
			name:     "happy number 7",
			n:        7,
			expected: true,
		},
		{
			name:     "happy number 1",
			n:        1,
			expected: true,
		},
		{
			name:     "unhappy number 2",
			n:        2,
			expected: false,
		},
		{
			name:     "unhappy number 4",
			n:        4,
			expected: false,
		},
		{
			name:     "unhappy number 16",
			n:        16,
			expected: false,
		},
		{
			name:     "happy number 10",
			n:        10,
			expected: true,
		},
		{
			name:     "happy number 13",
			n:        13,
			expected: true,
		},
		{
			name:     "unhappy number 20",
			n:        20,
			expected: false,
		},
		{
			name:     "large happy number 100",
			n:        100,
			expected: true,
		},
		{
			name:     "large unhappy number 999",
			n:        999,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHappy(tt.n)
			assert.Equal(t, tt.expected, result, "isHappy(%d) = %v, want %v", tt.n, result, tt.expected)
		})
	}
}