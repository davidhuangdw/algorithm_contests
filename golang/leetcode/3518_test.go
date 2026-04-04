package main

import "testing"

func Test_smallestPalindrome(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "abba", k: 2},
			want: "baab",
		},
		{
			name: "Example 2",
			args: args{s: "aa", k: 2},
			want: "",
		},
		{
			name: "Example 3",
			args: args{s: "bacab", k: 1},
			want: "abcba",
		},
		{
			name: "Length 1, k=1",
			args: args{s: "a", k: 1},
			want: "a",
		},
		{
			name: "Length 1, k=2",
			args: args{s: "a", k: 2},
			want: "",
		},
		{
			name: "All same chars",
			args: args{s: "aaaa", k: 1},
			want: "aaaa",
		},
		{
			name: "All same chars, k>1",
			args: args{s: "aaaa", k: 2},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestPalindrome(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("smallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
