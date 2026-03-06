package main

import (
	"testing"
)

func Test_convex2d(t *testing.T) {
	tests := []struct {
		name   string
		points [][2]float64
		want   [][2]float64
	}{
		{
			name:   "triangle",
			points: [][2]float64{{0, 0}, {1, 0}, {0, 1}},
			want:   [][2]float64{{0, 0}, {1, 0}, {0, 1}},
		},
		{
			name:   "square",
			points: [][2]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			want:   [][2]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
		},
		{
			name:   "square with interior point",
			points: [][2]float64{{0, 0}, {2, 0}, {2, 2}, {0, 2}, {1, 1}},
			want:   [][2]float64{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
		},
		{
			name:   "collinear points",
			points: [][2]float64{{0, 0}, {1, 1}, {2, 2}},
			want:   [][2]float64{{0, 0}, {2, 2}},
		},
		{
			name:   "pentagon",
			points: [][2]float64{{0, 0}, {2, 0}, {3, 2}, {1, 3}, {-1, 1}},
			want:   [][2]float64{{-1, 1}, {0, 0}, {2, 0}, {3, 2}, {1, 3}},
		},
		{
			name:   "all collinear horizontal",
			points: [][2]float64{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			want:   [][2]float64{{0, 0}, {3, 0}},
		},
		{
			name:   "all collinear vertical",
			points: [][2]float64{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			want:   [][2]float64{{0, 0}, {0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convex2d(tt.points)
			if len(got) != len(tt.want) {
				t.Errorf("convex2d() returned %v points, want %v points", len(got), len(tt.want))
				t.Errorf("got: %v", got)
				t.Errorf("want: %v", tt.want)
				return
			}
			for i := range got {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Errorf("convex2d()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
