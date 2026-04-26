package main

import "testing"

func Test_maxCollectedFruits(t *testing.T) {
	tests := []struct {
		name   string
		fruits [][]int
		want   int
	}{
		{
			name: "Example 1",
			fruits: [][]int{
				{1, 2, 3, 4},
				{5, 6, 8, 7},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: 100,
		},
		{
			name: "Small grid 2x2",
			fruits: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy fruits because the function might be destructive
			f := make([][]int, len(tt.fruits))
			for i := range tt.fruits {
				f[i] = make([]int, len(tt.fruits[i]))
				copy(f[i], tt.fruits[i])
			}
			if got := maxCollectedFruits(f); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
