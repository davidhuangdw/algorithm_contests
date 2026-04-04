package main

import (
	"testing"
)

func TestLargestPalindrome(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{3, 5, "595"},
		{1, 4, "8"},
		{5, 6, "89898"},
		{6, 9, "999999"},
		{7, 7, "9994999"},
		{2, 1, "99"},
	}

	for _, tt := range tests {
		got := largestPalindrome(tt.n, tt.k)
		if got != tt.want {
			t.Errorf("largestPalindrome(%d, %d) = %v; want %v", tt.n, tt.k, got, tt.want)
		}
	}
}
