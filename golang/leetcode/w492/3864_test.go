package main

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		encCost  int
		flatCost int
		want     int64
	}{
		{
			name:     "all zeros flat",
			s:        "0000",
			encCost:  10,
			flatCost: 5,
			want:     5,
		},
		{
			name:     "all ones split",
			s:        "1111",
			encCost:  1,
			flatCost: 100,
			want:     4,
		},
		{
			name:     "mixed split",
			s:        "1100",
			encCost:  2,
			flatCost: 10,
			want:     14,
		},
		{
			name:     "simple case",
			s:        "10",
			encCost:  5,
			flatCost: 3,
			want:     8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost(tt.s, tt.encCost, tt.flatCost)
			if got != tt.want {
				t.Errorf("minCost(%q, %v, %v) = %v, want %v", tt.s, tt.encCost, tt.flatCost, got, tt.want)
			}
		})
	}
}
