package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		flights  [][]int
		src      int
		dst      int
		k        int
		expected int
	}{
		{
			name: "LeetCode example 1 - direct flight with k=0",
			n:    3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:      0,
			dst:      2,
			k:        0,
			expected: 500,
		},
		{
			name: "LeetCode example 2 - one stop with k=1",
			n:    3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:      0,
			dst:      2,
			k:        1,
			expected: 200,
		},
		{
			name: "Same source and destination",
			n:    2,
			flights: [][]int{
				{0, 1, 50},
			},
			src:      0,
			dst:      0,
			k:        0,
			expected: 0,
		},
		{
			name:     "No flights available",
			n:        3,
			flights:  [][]int{},
			src:      0,
			dst:      2,
			k:        2,
			expected: -1,
		},
		{
			name: "Destination unreachable within k stops",
			n:    4,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{2, 3, 100},
			},
			src:      0,
			dst:      3,
			k:        1,
			expected: -1,
		},
		{
			name: "Multiple paths with different costs",
			n:    4,
			flights: [][]int{
				{0, 1, 1},
				{0, 2, 5},
				{1, 3, 1},
				{2, 3, 1},
			},
			src:      0,
			dst:      3,
			k:        2,
			expected: 2,
		},
		{
			name: "Complex network with optimal path requiring k stops",
			n:    5,
			flights: [][]int{
				{0, 1, 10},
				{0, 2, 1},
				{2, 1, 1},
				{1, 3, 1},
				{2, 4, 1},
				{3, 4, 1},
			},
			src:      0,
			dst:      3,
			k:        3,
			expected: 3,
		},
		{
			name: "Negative k value (invalid input)",
			n:    2,
			flights: [][]int{
				{0, 1, 100},
			},
			src:      0,
			dst:      1,
			k:        -1,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
