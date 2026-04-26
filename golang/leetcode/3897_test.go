package main

import "testing"

func Test_maxValue(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums0 []int
		want  int
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 1},
			nums0: []int{1, 1},
			want:  10, // "1010"
		},
		{
			name:  "Example 2",
			nums1: []int{2, 1},
			nums0: []int{1, 2},
			want:  52, // "110100"
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.nums1, tt.nums0); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
