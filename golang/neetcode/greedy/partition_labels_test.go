package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []int
	}{
		{
			name:     "leetcode example 1",
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
		{
			name:     "single character",
			s:        "a",
			expected: []int{1},
		},
		{
			name:     "all same characters",
			s:        "aaaaa",
			expected: []int{5},
		},
		{
			name:     "consecutive partitions",
			s:        "abcdefg",
			expected: []int{1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:     "overlapping partitions",
			s:        "ababacaca",
			expected: []int{9},
		},
		{
			name:     "two partitions",
			s:        "abcdeffedcba",
			expected: []int{12},
		},
		{
			name:     "multiple partitions with overlaps",
			s:        "abacabadabacaba",
			expected: []int{15},
		},
		{
			name:     "empty string",
			s:        "",
			expected: []int{},
		},
		{
			name:     "leetcode example 2",
			s:        "eccbbbbdec",
			expected: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partitionLabels(tt.s)
			assert.Equal(t, tt.expected, result, "partitionLabels(%q) = %v, want %v", tt.s, result, tt.expected)
		})
	}
}