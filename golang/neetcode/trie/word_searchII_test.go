package trie

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name:  "empty board",
			board: [][]byte{},
			words: []string{"test"},
			want:  []string{},
		},
		{
			name:  "empty words list",
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{},
			want:  []string{},
		},
		{
			name:  "basic word found",
			board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"oath", "eat"},
		},
		{
			name:  "single character board",
			board: [][]byte{{'a'}},
			words: []string{"a", "b", "c"},
			want:  []string{"a"},
		},
		{
			name:  "words with overlapping paths",
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{"ab", "ac", "ad", "bd", "cd"},
			want:  []string{"ab", "ac", "bd", "cd"},
		},
		{
			name:  "word not found due to visited cells",
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{"abc", "acb", "adb"},
			want:  []string{},
		},
		{
			name:  "duplicate words in result",
			board: [][]byte{{'a', 'a'}, {'a', 'a'}},
			words: []string{"a", "aa", "aaa", "aaaa"},
			want:  []string{"a", "aa", "aaa", "aaaa"},
		},
		{
			name:  "case sensitivity",
			board: [][]byte{{'A', 'B'}, {'C', 'D'}},
			words: []string{"AB", "ab", "CD", "cd"},
			want:  []string{"AB", "CD"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findWords(tt.board, tt.words)

			// Sort both slices for consistent comparison
			sort.Strings(result)
			sort.Strings(tt.want)

			assert.Equal(t, tt.want, result,
				"findWords() = %v, want %v", result, tt.want)
		})
	}
}

func TestFindWordsEdgeCases(t *testing.T) {
	t.Run("nil board", func(t *testing.T) {
		result := findWords(nil, []string{"test"})
		assert.Empty(t, result)
	})

	t.Run("nil words", func(t *testing.T) {
		result := findWords([][]byte{{'a'}}, nil)
		assert.Empty(t, result)
	})

	t.Run("single row board", func(t *testing.T) {
		board := [][]byte{{'h', 'e', 'l', 'l', 'o'}}
		words := []string{"hello", "world", "hell"}
		result := findWords(board, words)
		assert.ElementsMatch(t, []string{"hello", "hell"}, result)
	})

	t.Run("single column board", func(t *testing.T) {
		board := [][]byte{{'h'}, {'e'}, {'l'}, {'l'}, {'o'}}
		words := []string{"hello", "world", "hell"}
		result := findWords(board, words)
		assert.ElementsMatch(t, []string{"hello", "hell"}, result)
	})
}
