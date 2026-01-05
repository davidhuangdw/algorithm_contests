package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinExtraChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		dict []string
		want int
	}{
		{
			name: "empty string",
			s:    "",
			dict: []string{"a", "b"},
			want: 0,
		},
		{
			name: "string fully covered by dictionary",
			s:    "leetcode",
			dict: []string{"leet", "code"},
			want: 0,
		},
		{
			name: "string with one extra character",
			s:    "leetcodxe",
			dict: []string{"leet", "code"},
			want: 5,
		},
		{
			name: "string with multiple extra characters",
			s:    "sayhelloworld",
			dict: []string{"hello", "world"},
			want: 3,
		},
		{
			name: "single character string in dictionary",
			s:    "a",
			dict: []string{"a"},
			want: 0,
		},
		{
			name: "single character string not in dictionary",
			s:    "b",
			dict: []string{"a"},
			want: 1,
		},
		{
			name: "overlapping dictionary words",
			s:    "applepenapple",
			dict: []string{"apple", "pen"},
			want: 0,
		},
		{
			name: "empty dictionary",
			s:    "hello",
			dict: []string{},
			want: 5,
		},
		{
			name: "dictionary with single character words",
			s:    "abc",
			dict: []string{"a", "b", "c"},
			want: 0,
		},
		{
			name: "dictionary with longer words",
			s:    "abcdef",
			dict: []string{"abc", "def"},
			want: 0,
		},
		{
			name: "partial coverage with gaps",
			s:    "xabcxdefx",
			dict: []string{"abc", "def"},
			want: 3,
		},
		{
			name: "dictionary words with common prefixes",
			s:    "carcatcard",
			dict: []string{"car", "cat", "card"},
			want: 0,
		},
		{
			name: "case sensitivity test",
			s:    "HelloWorld",
			dict: []string{"hello", "world"},
			want: 10,
		},
		{
			name: "duplicate dictionary words",
			s:    "testtest",
			dict: []string{"test", "test"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minExtraChar(tt.s, tt.dict)
			assert.Equal(t, tt.want, result,
				"minExtraChar(%q, %v) = %d, want %d", tt.s, tt.dict, result, tt.want)
		})
	}
}

func TestMinExtraCharEdgeCases(t *testing.T) {
	t.Run("nil dictionary", func(t *testing.T) {
		result := minExtraChar("test", nil)
		assert.Equal(t, 4, result)
	})

	t.Run("very long string", func(t *testing.T) {
		s := "abcdefghijklmnopqrstuvwxyz"
		dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}
		result := minExtraChar(s, dict)
		assert.Equal(t, 0, result)
	})

	t.Run("dictionary with empty strings", func(t *testing.T) {
		result := minExtraChar("test", []string{"", "test"})
		assert.Equal(t, 0, result)
	})
}
