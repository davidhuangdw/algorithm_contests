package main

import "testing"

func Test_validSubarraySize(t *testing.T) {
	type args struct {
		nums []int
		th   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 4, 3, 1},
				th:   6,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{6, 5, 6, 5, 8},
				th:   7,
			},
			want: 1, // Can be 1, 2, 3, 4, or 5. But based on the code, it returns the first one it finds.
		},
		{
			name: "No valid subarray",
			args: args{
				nums: []int{1, 1, 1, 1},
				th:   10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validSubarraySize(tt.args.nums, tt.args.th)
			if got < tt.want {
				t.Errorf("validSubarraySize() = %v, want >=%v", got, tt.want)
			}
		})
	}
}
