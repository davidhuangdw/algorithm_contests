package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		// Basic cases
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AAAA", 2, 4},
		{"ABBB", 2, 4},
		{"ABCDE", 0, 1},

		// Edge cases
		{"", 0, 0},
		{"", 5, 0},
		{"A", 0, 1},
		{"A", 5, 1},
		{"AA", 0, 2},
		{"AA", 1, 2},
		{"AB", 0, 1},
		{"AB", 1, 2},
		{"ABA", 1, 3},
		{"AAB", 0, 2},

		// Complex cases
		{"AABABBA", 1, 4},
		{"ABABABAB", 2, 5},
		{"ABCDEFGH", 3, 4},
		{"AABBCCDD", 2, 4},
		{"AAAAABBB", 2, 7},
		{"ABABABAB", 0, 1},
		{"AABBAABB", 1, 3},
		{"ABCDABCD", 4, 6},
		{"A", 10, 1},
		{"AB", 10, 2},

		// Large k values
		{"ABCDE", 10, 5},
		{"AABBBCC", 5, 7},
		{"XYZXYZ", 10, 6},

		// All same character
		{"AAAAA", 0, 5},
		{"AAAAA", 1, 5},
		{"AAAAA", 5, 5},

		// Mixed characters
		{"AABACCD", 1, 4},
		{"AABACCDD", 2, 5},
		{"ABCABCABC", 2, 4},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			result := characterReplacement(tt.s, tt.k)
			assert.Equal(t, tt.want, result,
				"characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, result, tt.want)
		})
	}
}
