package leetcode

import "testing"

func Test_kthDistinct(t *testing.T) {
	type args struct {
		arr []string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]string{"d","b","c","b","c","a"}, 2}, "a"},
		{"2", args{[]string{"aaa","aa","a"}, 1}, "aaa"},
		{"3", args{[]string{"a","b","a"}, 3}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthDistinct(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kthDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
