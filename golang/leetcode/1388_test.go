package main

import "testing"

func Test_maxSizeSlices(t *testing.T) {
	tests := []struct {
		name   string
		slices []int
		want   int
	}{
		{
			name:   "Example 1",
			slices: []int{1, 2, 3, 4, 5, 6},
			want:   10,
		},
		{
			name:   "Example 2",
			slices: []int{8, 9, 8, 6, 1, 1},
			want:   16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSizeSlices(tt.slices); got != tt.want {
				t.Errorf("maxSizeSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
