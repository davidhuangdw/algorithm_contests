package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReach(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		minJump  int
		maxJump  int
		expected bool
	}{
		{
			name:     "LeetCode example 1",
			s:        "011010",
			minJump:  2,
			maxJump:  3,
			expected: true,
		},
		{
			name:     "LeetCode example 2",
			s:        "01101110",
			minJump:  2,
			maxJump:  3,
			expected: false,
		},
		{
			name:     "single character '0'",
			s:        "0",
			minJump:  1,
			maxJump:  1,
			expected: true,
		},
		{
			name:     "single character '1'",
			s:        "1",
			minJump:  1,
			maxJump:  1,
			expected: false,
		},
		{
			name:     "impossible due to first character '1'",
			s:        "1010",
			minJump:  1,
			maxJump:  2,
			expected: false,
		},
		{
			name:     "exact jump to end",
			s:        "0000",
			minJump:  3,
			maxJump:  3,
			expected: true,
		},
		{
			name:     "jump range too small",
			s:        "0000",
			minJump:  4,
			maxJump:  4,
			expected: false,
		},
		{
			name:     "multiple paths available",
			s:        "0000000",
			minJump:  2,
			maxJump:  4,
			expected: true,
		},
		{
			name:     "blocked by '1' in middle",
			s:        "0011100",
			minJump:  2,
			maxJump:  3,
			expected: false,
		},
		{
			name:     "large jump range",
			s:        "0000000000",
			minJump:  1,
			maxJump:  9,
			expected: true,
		},
		{
			name:     "minJump equals maxJump",
			s:        "00000",
			minJump:  2,
			maxJump:  2,
			expected: true,
		},
		{
			name:     "jump over obstacles",
			s:        "010010",
			minJump:  2,
			maxJump:  3,
			expected: true,
		},
		{
			name:     "cannot jump over wide obstacle",
			s:        "011110",
			minJump:  2,
			maxJump:  3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canReach(tt.s, tt.minJump, tt.maxJump)
			assert.Equal(t, tt.expected, result, "canReach(%q, %d, %d)", tt.s, tt.minJump, tt.maxJump)
		})
	}
}
