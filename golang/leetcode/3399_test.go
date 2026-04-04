package main

import "testing"

func Test_minLength(t *testing.T) {
	type args struct {
		s      string
		numOps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "000001", numOps: 1},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{s: "0000", numOps: 2},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{s: "0101", numOps: 0},
			want: 1,
		},
		{
			name: "All same, zero ops",
			args: args{s: "000", numOps: 0},
			want: 3,
		},
		{
			name: "All same, enough ops to break",
			args: args{s: "000000", numOps: 2},
			want: 2, // 001001 -> max len 2
		},
		{
			name: "Alternating with ops",
			args: args{s: "111000", numOps: 2},
			want: 1, // 111000 -> 101010 (2 flips: s[1], s[4])
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minLength(tt.args.s, tt.args.numOps); got != tt.want {
				t.Errorf("minLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
