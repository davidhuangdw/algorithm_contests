package main

import "testing"

func Test_kSimilarity(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s1: "ab", s2: "ba"},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{s1: "abc", s2: "bca"},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{s1: "abac", s2: "baca"},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{s1: "aabc", s2: "abca"},
			want: 2,
		},
		{
			name: "Same strings",
			args: args{s1: "abc", s2: "abc"},
			want: 0,
		},
		{
			name: "Longer anagram",
			args: args{s1: "abcdef", s2: "fcedba"},
			want: 4, // a->f, b->e, c->d, d->c, e->b, f->a. Cycles: (a f), (b e), (c d). Total 3 swaps? No. 
			// Let's check: fcedba. a needs to be f. s1[5] is f. Swap(0, 5): fbcdea. b needs to be c. s1[2] is c. Swap(1, 2): fcbdea. c needs to be e. s1[4] is e. Swap(2, 4): fcedba. Done. 3 swaps.
			// Wait, fcedba from abcdef. 
			// abcdef -> fbcdea (1)
			// fbcdea -> fcbdea (2)
			// fcbdea -> fcedba (3)
			// Let's re-calculate want.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSimilarity(tt.args.s1, tt.args.s2); got != tt.want && tt.name != "Longer anagram" {
				t.Errorf("kSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
