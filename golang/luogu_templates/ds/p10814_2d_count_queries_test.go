package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_count2DPoints(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][3]int
		want    []int
	}{
		{
			name:    "basic query - count points where value <= 2",
			nums:    []int{0, 3, 1, 2, 4},
			queries: [][3]int{{1, 3, 2}},
			want:    []int{2}, // positions 2,3 have values 1,2 <= 2
		},
		{
			name:    "multiple queries",
			nums:    []int{0, 3, 1, 2, 4, 5},
			queries: [][3]int{{1, 3, 2}, {2, 4, 3}, {1, 5, 5}},
			want:    []int{2, 2, 5},
		},
		{
			name:    "no matches - all values too large",
			nums:    []int{0, 5, 6, 7},
			queries: [][3]int{{1, 3, 3}},
			want:    []int{0},
		},
		{
			name:    "all matches",
			nums:    []int{0, 1, 2, 3},
			queries: [][3]int{{1, 3, 3}},
			want:    []int{3},
		},
		{
			name:    "single element queries",
			nums:    []int{0, 5},
			queries: [][3]int{{1, 1, 5}, {1, 1, 4}},
			want:    []int{1, 0},
		},
		{
			name:    "query from position 1",
			nums:    []int{0, 2, 4, 1, 3},
			queries: [][3]int{{1, 4, 2}},
			want:    []int{2}, // positions 1,3 have values 2,1 <= 2
		},
		{
			name:    "overlapping range queries",
			nums:    []int{0, 5, 3, 1, 4, 2},
			queries: [][3]int{{1, 3, 3}, {2, 5, 3}, {1, 5, 4}},
			want:    []int{2, 3, 4}, // [3,1], [3,1,2], [3,1,4,2]
		},
		{
			name:    "boundary value query",
			nums:    []int{0, 1, 1, 1, 1},
			queries: [][3]int{{1, 4, 1}, {1, 4, 0}},
			want:    []int{4, 0},
		},
		{
			name:    "large values with small threshold",
			nums:    []int{0, 100, 200, 50, 150},
			queries: [][3]int{{1, 4, 100}},
			want:    []int{2}, // positions 1,3 have values 100,50 <= 100
		},
		{
			name:    "duplicate values",
			nums:    []int{0, 3, 3, 3, 1, 2},
			queries: [][3]int{{1, 5, 3}, {2, 4, 2}},
			want:    []int{5, 1}, // all 5 values <= 3; position 4 has value 1 <= 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := count2DPoints(tt.nums, tt.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
