package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a", "b"},
			want:     true,
		},
		{
			name:     "single word match",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "single word exact match",
			s:        "apple",
			wordDict: []string{"apple"},
			want:     true,
		},
		{
			name:     "no match",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "multiple words match",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "repeated words",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			want:     true,
		},
		{
			name:     "overlapping words",
			s:        "catsand",
			wordDict: []string{"cat", "cats", "and", "sand"},
			want:     true,
		},
		{
			name:     "word longer than string",
			s:        "a",
			wordDict: []string{"aa"},
			want:     false,
		},
		{
			name:     "empty dictionary",
			s:        "test",
			wordDict: []string{},
			want:     false,
		},
		{
			name:     "single character match",
			s:        "a",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "complex overlapping",
			s:        "abcd",
			wordDict: []string{"a", "abc", "b", "cd"},
			want:     true,
		},
		{
			name:     "leetcode example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "leetcode example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "leetcode example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			assert.Equal(t, tt.want, got, "wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
		})
	}
}
