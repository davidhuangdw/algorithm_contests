package main

import "testing"

func Test_countTriangle(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][2]int
		want  int
	}{
		{
			name:  "simple triangle",
			n:     3,
			edges: [][2]int{{1, 2}, {2, 3}, {1, 3}},
			want:  1,
		},
		{
			name:  "complete graph K4",
			n:     4,
			edges: [][2]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			want:  4,
		},
		{
			name:  "no triangle",
			n:     4,
			edges: [][2]int{{1, 2}, {2, 3}, {3, 4}},
			want:  0,
		},
		{
			name:  "two triangles",
			n:     5,
			edges: [][2]int{{1, 2}, {2, 3}, {1, 3}, {3, 4}, {4, 5}, {3, 5}},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriangle(tt.n, tt.edges); got != tt.want {
				t.Errorf("countTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
