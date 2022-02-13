package leetcode

import "testing"

func Test_minimizedMaximum(t *testing.T) {
	type args struct {
		n          int
		quantities []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{6, []int{11, 6}}, 3},
		{"2", args{7, []int{15,10,10}}, 5},
		{"3", args{1, []int{100000}}, 100000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizedMaximum(tt.args.n, tt.args.quantities); got != tt.want {
				t.Errorf("minimizedMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
