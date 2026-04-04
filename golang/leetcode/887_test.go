package main

import "testing"

func Test_superEggDrop(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{k: 1, n: 2},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{k: 2, n: 6},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{k: 3, n: 14},
			want: 4,
		},
		{
			name: "k=2, n=1",
			args: args{k: 2, n: 1},
			want: 1,
		},
		{
			name: "k=1, n=1",
			args: args{k: 1, n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
