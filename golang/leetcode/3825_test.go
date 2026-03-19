package main

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
		// Additional comment for clarity
		comment string
	}{
		{
			name:    "no common bits",
			args:    args{[]int{0, 1, 12}},
			want:    1,
			comment: "Numbers have no common set bits, so each can only form subsequence of length 1",
		},
		{
			name:    "common bit with mixed order",
			args:    args{[]int{5, 4, 7}},
			want:    2,
			comment: "Numbers share bit 2 (0b100), longest increasing is [4,7] (length 2)",
		},
		{
			name:    "perfect increasing sequence",
			args:    args{[]int{2, 3, 6}},
			want:    3,
			comment: "Numbers share bit 1 (0b10), and all are strictly increasing",
		},
		{
			name:    "multiple common bits",
			args:    args{[]int{6, 8, 38, 40, 48, 28, 4}},
			want:    3,
			comment: "Numbers 38,40,48 share bit 5 (0b100000) and form increasing subsequence",
		},
		{
			name:    "single element",
			args:    args{[]int{5}},
			want:    1,
			comment: "Single element can only form subsequence of length 1",
		},
		{
			name:    "all elements same",
			args:    args{[]int{4, 4, 4}},
			want:    1,
			comment: "Elements are same, cannot form strictly increasing subsequence longer than 1",
		},
		{
			name:    "strictly increasing with no common bits",
			args:    args{[]int{1, 2, 4, 8}},
			want:    1,
			comment: "Each number has unique set bit, so no common bits",
		},
		{
			name:    "two different common bits groups",
			args:    args{[]int{1, 3, 5, 2, 4, 6}},
			want:    3,
			comment: "Two groups: odds (bit 0) have [1,3,5] (length 3), evens (bit 1) have [2,4,6] (length 3)",
		},
		{
			name:    "large numbers with shared bits",
			args:    args{[]int{1000, 2000, 3000, 1500}},
			want:    3,
			comment: "Numbers share multiple bits but longest increasing subsequence for any bit group is 3",
		},
		{
			name:    "decreasing sequence",
			args:    args{[]int{5, 3, 1}},
			want:    1,
			comment: "Sequence is strictly decreasing, no increasing subsequence longer than 1",
		},
		{
			name:    "with zero",
			args:    args{[]int{0, 2, 4, 6}},
			want:    2,
			comment: "Zero has no bits set, numbers 2(0b10) and 6(0b110) share bit 1, numbers 4(0b100) and 6(0b110) share bit 2",
		},
		{
			name:    "complex pattern",
			args:    args{[]int{9, 12, 5, 13, 14, 3, 10}},
			want:    4,
			comment: "Numbers 9,12,13,14 share bit 8 (0b10000), forming increasing subsequence [9,12,13,14]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.nums); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v\nInput: %v\nComment: %s", got, tt.want, tt.args.nums, tt.comment)
			}
		})
	}
}