package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// Basic cases
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
		{"abcde", 5},

		// Edge cases
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
		{"aba", 2},
		{"abba", 2},
		{"tmmzuxt", 5},
		{"abcabcbb", 3},

		// Complex cases
		{"abcdefghijklmnopqrstuvwxyz", 26},
		{"abcdefghijklmnopqrstuvwxyza", 26},
		{"abcdeabcde", 5},
		{"aabbccddeeff", 2},
		{"abcabcbbdefgh", 6},
		{"pwwkewabc", 6},
		{"abacadaeaf", 3},
		{"xyzxyzxyz", 3},
		{"abcdefgabcdefg", 7},

		// Unicode and special characters
		{"世界世界", 2},
		{"hello 世界", 5},
		{"a b c d e", 3},
		{"!@#$%^&*()", 10},
		{"a1b2c3d4e5", 10},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.s)
			assert.Equal(t, tt.want, result,
				"lengthOfLongestSubstring(%q) = %d, want %d", tt.s, result, tt.want)
		})
	}
}
