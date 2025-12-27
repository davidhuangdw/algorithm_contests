package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name:     "empty word should return true",
			board:    [][]byte{{'A', 'B'}, {'C', 'D'}},
			word:     "",
			expected: true,
		},
		{
			name:     "empty board should return false",
			board:    [][]byte{},
			word:     "A",
			expected: false,
		},
		{
			name: "basic case - word found horizontally",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "basic case - word found vertically",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "word not found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "single character board - word found",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "single character board - word not found",
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
		{
			name: "backtracking required - avoid reuse",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABD",
			expected: false, // Can't reuse B
		},
		{
			name: "complex path with backtracking",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			word:     "ABEDH",
			expected: true,
		},
		{
			name: "word longer than board area",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABCDE",
			expected: false,
		},
		{
			name: "duplicate characters - valid path",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AAAAAAAAA",
			expected: true,
		},
		{
			name: "case sensitivity - lowercase not found",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ab",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exist(tt.board, tt.word)
			assert.Equal(t, tt.expected, result, 
				"exist(%v, %s) should be %v", tt.board, tt.word, tt.expected)
		})
	}
}