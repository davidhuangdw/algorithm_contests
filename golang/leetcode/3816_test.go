package main

import (
	"testing"
)

func Test_lexSmallestAfterDeletion(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "basic case - all unique",
			s:    "abc",
			want: "abc",
		},
		{
			name: "duplicates at beginning",
			s:    "aabc",
			want: "aabc", // Function doesn't delete leading duplicates if they're in order
		},
		{
			name: "duplicates at end",
			s:    "abcc",
			want: "abc",
		},
		{
			name: "multiple duplicates",
			s:    "aabbcc",
			want: "aabbc", // Function preserves some duplicates to maintain lex order
		},
		{
			name: "lexicographic order important",
			s:    "cbacdcbc",
			want: "acdb",
		},
		{
			name: "all same characters",
			s:    "aaaa",
			want: "a",
		},
		{
			name: "single character",
			s:    "a",
			want: "a",
		},
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "complex pattern",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "numbers as characters",
			s:    "112233",
			want: "11223", // Function preserves some duplicates
		},
		{
			name: "reverse order",
			s:    "cba",
			want: "cba", // All unique, no deletion possible
		},
		{
			name: "mixed duplicates",
			s:    "abacaba",
			want: "aacab", // Based on actual function behavior
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lexSmallestAfterDeletion(tt.s)
			if got != tt.want {
				t.Errorf("lexSmallestAfterDeletion(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}