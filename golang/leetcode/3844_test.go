package main

import "testing"

func Test_almostPalindromic(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abacaba", 7},
		{"abcdef", 3},
		{"a", 1},
		{"aa", 2},
		{"abc", 3},
		{"racecar", 7},
		{"abcdcba", 7},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := almostPalindromic(tt.s); got != tt.want {
				t.Errorf("almostPalindromic(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
