package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "LeetCode example 1 - valid transformation",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			name:      "LeetCode example 2 - no transformation path",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "Same begin and end word",
			beginWord: "test",
			endWord:   "test",
			wordList:  []string{"best", "rest"},
			expected:  1,
		},
		{
			name:      "End word not in word list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "Direct transformation",
			beginWord: "hot",
			endWord:   "dot",
			wordList:  []string{"dot", "lot", "dog"},
			expected:  2,
		},
		{
			name:      "Begin word not in word list",
			beginWord: "abc",
			endWord:   "abd",
			wordList:  []string{"abd", "abe", "abf"},
			expected:  2,
		},
		{
			name:      "Multiple paths - shortest should be found",
			beginWord: "aaa",
			endWord:   "bbb",
			wordList:  []string{"aab", "abb", "aba", "baa", "bbb"},
			expected:  4,
		},
		{
			name:      "Single character difference required",
			beginWord: "abc",
			endWord:   "def",
			wordList:  []string{"abd", "abf", "aef", "def"},
			expected:  4,
		},
		{
			name:      "Complex word ladder",
			beginWord: "game",
			endWord:   "thee",
			wordList:  []string{"fame", "fate", "gate", "hate", "have", "hive", "give", "live"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			assert.Equal(t, tt.expected, result)
		})
	}
}
