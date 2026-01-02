package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name     string
		times    [][]int
		n        int
		k        int
		expected int
	}{
		{
			name:     "LeetCode example 1 - all nodes reachable",
			times:    [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:        4,
			k:        2,
			expected: 2,
		},
		{
			name:     "LeetCode example 2 - single node",
			times:    [][]int{{1, 2, 1}},
			n:        2,
			k:        1,
			expected: 1,
		},
		{
			name:     "LeetCode example 3 - unreachable nodes",
			times:    [][]int{{1, 2, 1}},
			n:        2,
			k:        2,
			expected: -1,
		},
		{
			name:     "Multiple paths to same node",
			times:    [][]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 1}, {3, 4, 1}},
			n:        4,
			k:        1,
			expected: 3,
		},
		{
			name:     "Cycle with shortest path",
			times:    [][]int{{1, 2, 1}, {2, 3, 1}, {3, 1, 1}, {3, 4, 1}},
			n:        4,
			k:        1,
			expected: 3,
		},
		{
			name:     "Disconnected graph - some nodes unreachable",
			times:    [][]int{{1, 2, 1}, {3, 4, 1}},
			n:        4,
			k:        1,
			expected: -1,
		},
		{
			name:     "Single node graph",
			times:    [][]int{},
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name:     "Multiple edges with different weights",
			times:    [][]int{{1, 2, 2}, {1, 3, 1}, {2, 4, 1}, {3, 4, 3}, {3, 2, 1}},
			n:        4,
			k:        1,
			expected: 3,
		},
		{
			name:     "Large weights and multiple paths",
			times:    [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 1}, {2, 4, 1}},
			n:        4,
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := networkDelayTime(tt.times, tt.n, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
