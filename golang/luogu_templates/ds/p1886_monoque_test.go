package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowMin(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		k       int
		less    func(a, b int) bool
		want    []int
		wantSub []int // Expected result from index k-1 onwards (full window results)
	}{
		{
			name:    "empty array",
			arr:     []int{},
			k:       1,
			less:    func(a, b int) bool { return a < b },
			want:    []int{},
			wantSub: []int{},
		},
		{
			name:    "k=1 (single element window)",
			arr:     []int{3, 1, 4, 1, 5, 9},
			k:       1,
			less:    func(a, b int) bool { return a < b },
			want:    []int{3, 1, 4, 1, 5, 9},
			wantSub: []int{3, 1, 4, 1, 5, 9},
		},
		{
			name:    "k=2, increasing",
			arr:     []int{1, 2, 3, 4, 5},
			k:       2,
			less:    func(a, b int) bool { return a < b },
			want:    []int{1, 1, 2, 3, 4},
			wantSub: []int{1, 2, 3, 4},
		},
		{
			name:    "k=2, decreasing",
			arr:     []int{5, 4, 3, 2, 1},
			k:       2,
			less:    func(a, b int) bool { return a < b },
			want:    []int{5, 4, 3, 2, 1},
			wantSub: []int{4, 3, 2, 1},
		},
		{
			name:    "k=3, mixed",
			arr:     []int{3, 1, 4, 1, 5, 9, 2, 6},
			k:       3,
			less:    func(a, b int) bool { return a < b },
			want:    []int{3, 1, 1, 1, 1, 1, 2, 2},
			wantSub: []int{1, 1, 1, 1, 2, 2},
		},
		{
			name:    "k=len(arr)",
			arr:     []int{5, 3, 8, 1, 2},
			k:       5,
			less:    func(a, b int) bool { return a < b },
			want:    []int{5, 3, 3, 1, 1},
			wantSub: []int{1},
		},
		{
			name:    "using greater than for max",
			arr:     []int{3, 1, 4, 1, 5, 9},
			k:       3,
			less:    func(a, b int) bool { return a > b },
			want:    []int{3, 3, 4, 4, 5, 9},
			wantSub: []int{4, 4, 5, 9},
		},
		{
			name:    "all same elements",
			arr:     []int{2, 2, 2, 2, 2},
			k:       3,
			less:    func(a, b int) bool { return a < b },
			want:    []int{2, 2, 2, 2, 2},
			wantSub: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := windowMin(tt.arr, tt.k, tt.less)
			assert.Equal(t, tt.want, got, "full result should match")

			// Test the full window results (from k-1 onwards)
			if tt.k > 0 && len(got) >= tt.k {
				gotSub := got[tt.k-1:]
				assert.Equal(t, tt.wantSub, gotSub, "full window results should match")
			} else {
				assert.Equal(t, tt.wantSub, []int{}, "empty result expected for invalid k")
			}
		})
	}
}
