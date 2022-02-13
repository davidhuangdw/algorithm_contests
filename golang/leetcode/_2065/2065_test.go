package _2065

import "testing"

func Test_maximalPathQuality(t *testing.T) {
	type args struct {
		values  []int
		edges   [][]int
		maxTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 32, 10, 43}, [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}, 49}, 75},
		{"2", args{[]int{61,11,18,43,81}, [][]int{{0,3,45},{0,2,16},{0,1,36},{3,4,38},{2,3,29}}, 88}, 79},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalPathQuality(tt.args.values, tt.args.edges, tt.args.maxTime); got != tt.want {
				t.Errorf("maximalPathQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
