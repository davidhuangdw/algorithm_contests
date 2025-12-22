package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		ss   string
		tt   string
		want string
	}{
		// Basic cases
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"aa", "aa", "aa"},
		{"abc", "ac", "abc"},

		// Edge cases
		{"", "", ""},
		{"", "a", ""},
		{"a", "", ""},
		{"abc", "", ""},
		{"a", "b", ""},
		{"aa", "a", "a"},
		{"aaa", "aa", "aa"},
		{"ab", "a", "a"},
		{"ab", "b", "b"},

		// Complex cases
		{"this is a test string", "tist", "t stri"},
		{"geeksforgeeks", "ork", "ksfor"},
		{"donutsandwaffles", "flea", "affle"},
		{"whoopiepiesmakemyscalegroan", "roam", "myscalegro"},
		{"coffeeteabiscuits", "cake", ""},
		{"abcalgorithmtest", "got", "gorit"},
		{"abcdefghijklmnopqrstuvwxyz", "az", "abcdefghijklmnopqrstuvwxyz"},
		{"abcdefghijklmnopqrstuvwxyz", "za", "abcdefghijklmnopqrstuvwxyz"},

		// Multiple occurrences
		{"abacaba", "aba", "aba"},
		{"abacabad", "aba", "aba"},
		{"abacabadabacaba", "aba", "aba"},
		{"abcabac", "aba", "aba"},
		{"abcabac", "aab", "aba"},

		// Case sensitivity (assuming case-sensitive)
		{"ADOBECODEBANC", "abc", ""},
		{"adobecodebanc", "ABC", ""},
		{"AaBbCc", "ABC", "AaBbC"},
		{"AaBbCc", "abc", "aBbCc"},

		// Special characters
		{"hello world!", "!dl", "ld!"},
		{"test@email.com", "@.", "@email."},
		{"1234567890", "135", "12345"},
		{"a1b2c3d4", "123", "1b2c3"},

		// Long strings
		{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "az", "za"},
		{"aaaaaaaaaabbbbbbbbbbcccccccccc", "abc", "abbbbbbbbbbc"},
		{"abcabcabcabcabcabcabcabc", "abc", "abc"},
		{"xyzxyzxyzxyzxyzxyz", "xyz", "xyz"},

		// No solution cases
		{"abc", "d", ""},
		{"hello", "xyz", ""},
		{"test", "testtest", ""},
		{"short", "longerthanshort", ""},
		{"a", "aa", ""},
		{"ab", "abc", ""},
	}

	for _, tt := range tests {
		t.Run(tt.ss+"_contains_"+tt.tt, func(t *testing.T) {
			result := minWindow(tt.ss, tt.tt)
			assert.Equal(t, tt.want, result,
				"minWindow(%q, %q) = %q, want %q", tt.ss, tt.tt, result, tt.want)
		})
	}
}
