package backtrace

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     []string
	}{
		{
			name:     "basic case with multiple sentences",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cat sand dog", "cats and dog"},
		},
		{
			name:     "single word",
			s:        "a",
			wordDict: []string{"a"},
			want:     []string{"a"},
		},
		{
			name:     "no valid sentences",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     []string{},
		},
		{
			name:     "multiple word options",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a", "b"},
			want:     []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)

			// Sort both slices to handle ordering differences
			sort.Strings(got)
			sort.Strings(tt.want)

			assert.Equal(t, tt.want, got, "wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
		})
	}
}
