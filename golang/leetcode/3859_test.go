package main

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "basic: k=2 m=2",
			args: args{[]int{1, 2, 1, 2, 2}, 2, 2},
			want: 2, // [1,2,1,2], [1,2,1,2,2]
		},
		{
			name: "basic: k=2 m=1",
			args: args{[]int{3, 1, 2, 4}, 2, 1},
			want: 3, // [3,1], [1,2], [2,4]
		},
		{
			name: "all same value k=1 m=3",
			args: args{[]int{1, 1, 1, 1}, 1, 3},
			want: 3, // [1,1,1] starting at 0, [1,1,1,1], [1,1,1] starting at 1
		},
		{
			name: "k=1 m=1 all valid",
			args: args{[]int{5, 5, 5}, 1, 1},
			want: 6, // [5], [5,5], [5,5,5], [5], [5,5], [5]
		},
		{
			name: "impossible - not enough distinct",
			args: args{[]int{1, 2, 3}, 4, 1},
			want: 0, // need 4 distinct values but only have 3
		},
		{
			name: "large m=3 requirement",
			args: args{[]int{1, 1, 1, 2, 2, 2}, 2, 3},
			want: 1, // only [1,1,1,2,2,2]
		},
		{
			name: "single element",
			args: args{[]int{1}, 1, 1},
			want: 1, // [1]
		},
		{
			name: "k=3 m=2 exact match",
			args: args{[]int{1, 2, 3, 1, 2, 3}, 3, 2},
			want: 1, // [1,2,3,1,2,3]
		},
		{
			name: "alternating k=2 m=3",
			args: args{[]int{1, 2, 1, 2, 1, 2}, 2, 3},
			want: 1, // [1,2,1,2,1,2]
		},
		{
			name: "sliding windows",
			args: args{[]int{1, 1, 2, 2, 3, 3}, 2, 2},
			want: 2, // [1,1,2,2], [2,2,3,3]
		},
		{
			name: "k=1 three distinct elements",
			args: args{[]int{1, 2, 3}, 1, 1},
			want: 3, // [1], [2], [3]
		},
		{
			name: "exactly k=2 m=2",
			args: args{[]int{5, 5, 6, 6, 7}, 2, 2},
			want: 1, // [5,5,6,6]
		},
		{
			name: "no valid - need m=2 each",
			args: args{[]int{1, 2, 3, 4, 5}, 2, 2},
			want: 0, // no value appears ≥2 times
		},
		{
			name: "3x3 perfect grid",
			args: args{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3, 3},
			want: 1, // [1,1,1,2,2,2,3,3,3]
		},
		{
			name: "k equals n",
			args: args{[]int{1, 2, 3, 4}, 4, 1},
			want: 1, // [1,2,3,4]
		},
		{
			name: "two same k=1 m=2",
			args: args{[]int{7, 7}, 1, 2},
			want: 1, // [7,7]
		},
		{
			name: "interleaved - only one valid",
			args: args{[]int{1, 2, 1, 2, 3, 3}, 2, 2},
			want: 1, // [1,2,1,2] only; [3,3] has only 1 distinct
		},
		{
			name: "consecutive k=3 m=1",
			args: args{[]int{1, 2, 3, 4, 5}, 3, 1},
			want: 3, // [1,2,3], [2,3,4], [3,4,5]
		},
		{
			name: "sparse repeats",
			args: args{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 3, 3},
			want: 1, // entire array
		},
		{
			name: "zero elements edge",
			args: args{[]int{0, 0, 1, 1}, 2, 2},
			want: 1, // [0,0,1,1]
		},
		{
			name: "mixed frequencies",
			args: args{[]int{1, 1, 2, 2, 2, 3}, 2, 2},
			want: 2, // [1,1,2,2], [1,1,2,2,2]
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubarrays(tt.args.nums, tt.args.k, tt.args.m)
			if got != tt.want {
				t.Errorf("countSubarrays(%v, k=%v, m=%v) = %v, want %v",
					tt.args.nums, tt.args.k, tt.args.m, got, tt.want)
			}
		})
	}
}
