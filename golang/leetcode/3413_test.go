package main

import "testing"

func Test_maximumCoins(t *testing.T) {
	type args struct {
		coins [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				coins: [][]int{{8, 10, 1}, {1, 3, 2}, {5, 6, 4}},
				k:     4,
			},
			want: 10,
		},
		{
			name: "Example 2",
			args: args{
				coins: [][]int{{1, 1, 1000}, {7, 11, 2}},
				k:     2,
			},
			want: 1000,
		},
		{
			name: "Large Window",
			args: args{
				coins: [][]int{{1, 3, 2}, {5, 6, 4}},
				k:     10,
			},
			want: 14,
		},
		{
			name: "Small Window",
			args: args{
				coins: [][]int{{1, 10, 5}},
				k:     3,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCoins(tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("maximumCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
