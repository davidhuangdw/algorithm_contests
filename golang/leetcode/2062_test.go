package leetcode

import "testing"

func Test_countVowelSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"aeiouu"}, 2},
		{"2", args{"unicornarihan"}, 0},
		{"3", args{"cuaieuouac"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelSubstrings(tt.args.word); got != tt.want {
				t.Errorf("countVowelSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
