package main

import "testing"

func Test_sortableIntegers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Test Case 1", []int{1, 2, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortableIntegers(tt.nums); got != tt.want {
				t.Errorf("sortableIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
