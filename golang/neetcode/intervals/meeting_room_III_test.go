package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostBooked(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		meetings [][]int
		expected int
	}{
		{
			name:     "single room, single meeting",
			n:        1,
			meetings: [][]int{{0, 10}},
			expected: 0,
		},
		{
			name:     "multiple rooms, single meeting",
			n:        3,
			meetings: [][]int{{0, 10}},
			expected: 0,
		},
		{
			name:     "LeetCode example 1",
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			expected: 0,
		},
		{
			name:     "LeetCode example 2",
			n:        3,
			meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			expected: 1,
		},
		{
			name:     "all meetings in one room",
			n:        2,
			meetings: [][]int{{0, 5}, {5, 10}, {10, 15}},
			expected: 0,
		},
		{
			name:     "balanced room usage",
			n:        2,
			meetings: [][]int{{0, 5}, {0, 5}, {5, 10}, {5, 10}},
			expected: 0, // Room 0 gets first two meetings
		},
		{
			name:     "overlapping meetings with room delay",
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			expected: 0,
		},
		{
			name:     "large gap between meetings",
			n:        2,
			meetings: [][]int{{0, 5}, {100, 105}, {200, 205}},
			expected: 0,
		},
		{
			name:     "meetings requiring room delay",
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}, {5, 6}},
			expected: 1,
		},
		{
			name:     "tie in room usage",
			n:        3,
			meetings: [][]int{{0, 5}, {0, 5}, {0, 5}, {5, 10}, {5, 10}, {5, 10}},
			expected: 0, // Room 0 gets first meeting in each batch
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mostBooked(tt.n, tt.meetings)
			assert.Equal(t, tt.expected, result, "mostBooked(%d, %v)", tt.n, tt.meetings)
		})
	}
}
