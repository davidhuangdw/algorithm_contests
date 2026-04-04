package main

import (
	"testing"
)

func Test_maxBuilding(t *testing.T) {
	type args struct {
		n            int
		restrictions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: 5,
				restrictions: [][]int{{2, 1}, {4, 1}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				n: 6,
				restrictions: [][]int{},
			},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{
				n: 10,
				restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy restrictions to avoid modifying the original test case data
			// as maxBuilding modifies the input slice.
			restrs := make([][]int, len(tt.args.restrictions))
			for i, r := range tt.args.restrictions {
				restrs[i] = []int{r[0], r[1]}
			}
			if got := maxBuilding(tt.args.n, restrs); got != tt.want {
				t.Errorf("maxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
