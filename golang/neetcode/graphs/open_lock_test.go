package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenLock(t *testing.T) {
	tests := []struct {
		name     string
		deadends []string
		target   string
		expected int
	}{
		{
			name:     "target is starting position",
			deadends: []string{},
			target:   "0000",
			expected: 0,
		},
		{
			name:     "one step to target",
			deadends: []string{},
			target:   "0001",
			expected: 1,
		},
		{
			name:     "leetcode example 1",
			deadends: []string{"0201", "0101", "0102", "1212", "2002"},
			target:   "0202",
			expected: 6,
		},
		{
			name:     "leetcode example 2",
			deadends: []string{"8888"},
			target:   "0009",
			expected: 1,
		},
		{
			name:     "deadend blocks target",
			deadends: []string{"0001"},
			target:   "0001",
			expected: -1,
		},
		{
			name:     "deadend blocks all paths",
			deadends: []string{"0001", "1000", "0100", "0010"},
			target:   "2222",
			expected: 10,
		},
		{
			name:     "multiple steps no deadends",
			deadends: []string{},
			target:   "2222",
			expected: 8,
		},
		{
			name:     "target with different digits",
			deadends: []string{"1234"},
			target:   "5678",
			expected: 14,
		},
		{
			name:     "deadend at starting position",
			deadends: []string{"0000"},
			target:   "1234",
			expected: -1,
		},
		{
			name:     "target is blocked by deadend",
			deadends: []string{"0009", "0090", "0900", "9000"},
			target:   "0000",
			expected: 0,
		},
		{
			name:     "complex deadends pattern",
			deadends: []string{"1000", "9000", "0100", "0900", "0010", "0090", "0001", "0009"},
			target:   "2222",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := openLock(tt.deadends, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}
