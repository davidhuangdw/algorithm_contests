package main

import "testing"

func TestMaximumNumberOfOnes(t *testing.T) {
	tests := []struct {
		name       string
		width      int
		height     int
		sideLength int
		maxOnes    int
		want       int
	}{
		{
			name:       "Example 1",
			width:      3,
			height:     3,
			sideLength: 2,
			maxOnes:    1,
			want:       4,
		},
		{
			name:       "Example 2",
			width:      3,
			height:     3,
			sideLength: 2,
			maxOnes:    2,
			want:       6,
		},
		{
			name:       "Example 3",
			width:      10,
			height:     10,
			sideLength: 4,
			maxOnes:    5,
			want:       42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumNumberOfOnes(tt.width, tt.height, tt.sideLength, tt.maxOnes); got != tt.want {
				t.Errorf("maximumNumberOfOnes(%v, %v, %v, %v) = %v, want %v",
					tt.width, tt.height, tt.sideLength, tt.maxOnes, got, tt.want)
			}
		})
	}
}
