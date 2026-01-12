package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarPooling(t *testing.T) {
	tests := []struct {
		name     string
		trips    [][]int
		capacity int
		expected bool
	}{
		{
			name:     "LeetCode Example 1",
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 4,
			expected: false,
		},
		{
			name:     "LeetCode Example 2",
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 5,
			expected: true,
		},
		{
			name:     "Single trip within capacity",
			trips:    [][]int{{3, 1, 5}},
			capacity: 3,
			expected: true,
		},
		{
			name:     "Single trip exceeds capacity",
			trips:    [][]int{{5, 1, 5}},
			capacity: 4,
			expected: false,
		},
		{
			name:     "No trips",
			trips:    [][]int{},
			capacity: 5,
			expected: true,
		},
		{
			name:     "Overlapping trips within capacity",
			trips:    [][]int{{2, 1, 4}, {3, 2, 3}, {1, 3, 6}},
			capacity: 5,
			expected: true,
		},
		{
			name:     "Overlapping trips exceed capacity",
			trips:    [][]int{{3, 1, 4}, {4, 2, 5}, {2, 3, 6}},
			capacity: 6,
			expected: false,
		},
		{
			name:     "Back-to-back trips",
			trips:    [][]int{{2, 1, 3}, {3, 3, 5}, {1, 5, 7}},
			capacity: 4,
			expected: true,
		},
		{
			name:     "Same pickup and dropoff time",
			trips:    [][]int{{2, 1, 1}, {3, 1, 1}},
			capacity: 5,
			expected: true,
		},
		{
			name:     "Large capacity with many trips",
			trips:    [][]int{{10, 1, 10}, {20, 2, 9}, {15, 3, 8}, {5, 4, 7}},
			capacity: 50,
			expected: true,
		},
		{
			name:     "Exact capacity at peak",
			trips:    [][]int{{3, 1, 4}, {2, 2, 5}, {1, 3, 6}},
			capacity: 6,
			expected: true,
		},
		{
			name:     "Dropoff before pickup",
			trips:    [][]int{{3, 5, 10}, {2, 1, 4}},
			capacity: 5,
			expected: true,
		},
		{
			name:     "Zero capacity with no trips",
			trips:    [][]int{},
			capacity: 0,
			expected: true,
		},
		{
			name:     "Zero capacity with trips",
			trips:    [][]int{{1, 1, 2}},
			capacity: 0,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := carPooling(tt.trips, tt.capacity)
			assert.Equal(t, tt.expected, result)
		})
	}
}
