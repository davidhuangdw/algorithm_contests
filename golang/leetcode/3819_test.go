package main

import (
	"testing"
)

func Test_rotateElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic rotation",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "with negative numbers",
			nums: []int{-1, 2, -3, 4, -5, 6},
			k:    1,
			want: []int{-1, 4, -3, 6, -5, 2}, // Positive elements: 2,4,6 → rotate by 1: 4,6,2
		},
		{
			name: "all negative",
			nums: []int{-1, -2, -3},
			k:    2,
			want: []int{-1, -2, -3},
		},
		{
			name: "k larger than positive count",
			nums: []int{1, 2, 3, -4, -5},
			k:    5,
			want: []int{3, 1, 2, -4, -5}, // Positive elements: 1,2,3 → rotate by 5%3=2: 3,1,2
		},
		{
			name: "single positive element",
			nums: []int{5, -1, -2},
			k:    1,
			want: []int{5, -1, -2},
		},
		{
			name: "no positive elements",
			nums: []int{-1, -2, -3},
			k:    3,
			want: []int{-1, -2, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateElements(tt.nums, tt.k)
			if !equalSlices(got, tt.want) {
				t.Errorf("rotateElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}