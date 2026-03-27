package main

import (
	"testing"
)

func TestMaximumANDSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		numSlots int
		want     int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6},
			numSlots: 3,
			want:     9,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 10, 4, 7, 1},
			numSlots: 9,
			want:     24,
		},
		{
			name:     "Minimum values",
			nums:     []int{1},
			numSlots: 1,
			want:     1,
		},
		{
			name:     "Slots more than nums",
			nums:     []int{1, 1},
			numSlots: 2,
			want:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumANDSum(tt.nums, tt.numSlots); got != tt.want {
				t.Errorf("maximumANDSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
