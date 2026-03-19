package main

import "testing"

func Test_alternatingXOR(t *testing.T) {
	type args struct {
		nums    []int
		target1 int
		target2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 1, 4}, 1, 5}, 1},
		{"2", args{[]int{1, 0, 0}, 1, 0}, 3},
		{"single element target1", args{[]int{5}, 5, 0}, 1},      // Single element matching target1
		{"single element target2", args{[]int{5}, 0, 5}, 0},      // Single element matching target2
		{"two elements alternating", args{[]int{1, 2}, 1, 2}, 1}, // [1^2] = 3, but should alternate
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alternatingXOR(tt.args.nums, tt.args.target1, tt.args.target2)
			if got != tt.want {
				t.Errorf("alternatingXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
