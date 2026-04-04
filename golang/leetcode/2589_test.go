package main

import "testing"

func Test_findMinimumTime(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				tasks: [][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				tasks: [][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}},
			},
			want: 4,
		},
		{
			name: "Overlap at end",
			args: args{
				tasks: [][]int{{1, 10, 1}, {10, 11, 1}},
			},
			want: 1, // Both can use time 10
		},
		{
			name: "No overlap",
			args: args{
				tasks: [][]int{{1, 2, 2}, {4, 5, 2}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumTime(tt.args.tasks); got != tt.want {
				t.Errorf("findMinimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
