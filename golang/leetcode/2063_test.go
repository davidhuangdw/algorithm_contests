package leetcode

import "testing"

func Test_countVowels(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"aba", args{"aba"}, 6},
		{"abc", args{"abc"}, 3},
		{"ltcd", args{"ltcd"}, 0},
		{"noosabasboosa", args{"noosabasboosa"}, 237},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowels(tt.args.word); got != tt.want {
				t.Errorf("countVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
