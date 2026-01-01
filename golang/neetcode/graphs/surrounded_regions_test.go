package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected [][]byte
	}{
		{
			name: "LeetCode example 1",
			input: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "LeetCode example 2",
			input: [][]byte{
				{'X'},
			},
			expected: [][]byte{
				{'X'},
			},
		},
		{
			name: "All X's",
			input: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
		},
		{
			name: "All O's",
			input: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
		},
		{
			name: "Single O surrounded",
			input: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "O connected to border",
			input: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "Complex pattern",
			input: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
		},
		{
			name: "Empty board",
			input: [][]byte{
				{},
			},
			expected: [][]byte{
				{},
			},
		},
		{
			name: "1x1 O",
			input: [][]byte{
				{'O'},
			},
			expected: [][]byte{
				{'O'},
			},
		},
		{
			name: "Mixed regions",
			input: [][]byte{
				{'O', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'O'},
			},
			expected: [][]byte{
				{'O', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'O'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy of the input board
			boardCopy := make([][]byte, len(tt.input))
			for i := range tt.input {
				boardCopy[i] = make([]byte, len(tt.input[i]))
				copy(boardCopy[i], tt.input[i])
			}

			solve(boardCopy)

			assert.Equal(t, tt.expected, boardCopy)
		})
	}
}
