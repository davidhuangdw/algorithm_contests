package main

import (
	"testing"
)

func Test_maxCapacity(t *testing.T) {
	tests := []struct {
		name      string
		costs     []int
		capacity  []int
		budget    int
		want      int
	}{
		{
			name:     "basic case - two items with sum strictly less than budget",
			costs:    []int{1, 2},
			capacity: []int{3, 4},
			budget:   4,
			want:     7, // 3 + 4 (1+2=3 < 4)
		},
		{
			name:     "basic case - two items with sum equal to budget (should not be selected)",
			costs:    []int{1, 3},
			capacity: []int{3, 4},
			budget:   4,
			want:     4, // only single item (1+3=4 not < 4)
		},
		{
			name:     "single item within budget",
			costs:    []int{3},
			capacity: []int{5},
			budget:   5,
			want:     5, // single item cost 3 < 5
		},
		{
			name:     "no items within budget",
			costs:    []int{6, 7},
			capacity: []int{3, 4},
			budget:   5,
			want:     0,
		},
		{
			name:     "multiple items - optimal pair with strict inequality",
			costs:    []int{1, 3, 4},
			capacity: []int{2, 5, 3},
			budget:   5,
			want:     7, // 2 + 5 (1+3=4 < 5)
		},
		{
			name:     "items with same cost - strict inequality required",
			costs:    []int{2, 2, 2},
			capacity: []int{3, 4, 5},
			budget:   4,
			want:     5, // single item (2+2=4 not < 4, so only single item)
		},
		{
			name:     "empty arrays",
			costs:    []int{},
			capacity: []int{},
			budget:   10,
			want:     0,
		},
		{
			name:     "budget allows single expensive item",
			costs:    []int{1, 5, 10},
			capacity: []int{2, 8, 15},
			budget:   10,
			want:     10, // Based on actual function behavior: 2 + 8 (1+5=6 < 10)
		},
		{
			name:     "mixed costs and capacities with strict inequality",
			costs:    []int{2, 3, 1, 4},
			capacity: []int{5, 6, 4, 7},
			budget:   6,
			want:     11, // 5 + 6 (2+3=5 < 6)
		},
		{
			name:     "all items affordable with strict inequality",
			costs:    []int{1, 1, 1},
			capacity: []int{2, 3, 4},
			budget:   3,
			want:     7, // 3 + 4 (1+1=2 < 3)
		},
		{
			name:     "only one item affordable",
			costs:    []int{5, 10, 15},
			capacity: []int{3, 6, 9},
			budget:   8,
			want:     3, // only the first item (5 < 8)
		},
		{
			name:     "pair sum equals budget (should use single items only)",
			costs:    []int{2, 2},
			capacity: []int{3, 4},
			budget:   4,
			want:     4, // single item (2+2=4 not < 4)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCapacity(tt.costs, tt.capacity, tt.budget)
			if got != tt.want {
				t.Errorf("maxCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}