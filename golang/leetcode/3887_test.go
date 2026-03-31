package main

import "testing"

func Test_numberOfEdgesAdded(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"Test Case 1", 3, [][]int{{0, 1, 1}, {1, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfEdgesAdded(tt.n, tt.edges); got != tt.want {
				t.Errorf("numberOfEdgesAdded() = %v, want %v", got, tt.want)
			}
		})
	}
}
