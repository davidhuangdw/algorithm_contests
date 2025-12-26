package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		// Basic cases
		{"single task", []byte{'A'}, 0, 1},
		{"single task with cooldown", []byte{'A'}, 2, 1},
		{"two different tasks", []byte{'A', 'B'}, 0, 2},
		{"two different tasks with cooldown", []byte{'A', 'B'}, 1, 2},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{"LeetCode example 2", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{"LeetCode example 3", []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},

		// Edge cases
		{"all same tasks", []byte{'A', 'A', 'A'}, 1, 5},
		{"all same tasks with large cooldown", []byte{'A', 'A', 'A'}, 3, 9},
		{"all different tasks", []byte{'A', 'B', 'C', 'D'}, 1, 4},
		{"empty tasks", []byte{}, 2, 0},

		// Complex cases
		{"mixed frequency", []byte{'A', 'A', 'A', 'B', 'B', 'C', 'C'}, 1, 7},
		{"uneven distribution", []byte{'A', 'A', 'A', 'A', 'B', 'B', 'C'}, 2, 10},
		{"many tasks with cooldown", []byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leastInterval(tt.tasks, tt.n)
			assert.Equal(t, tt.expected, result,
				"leastInterval(%v, %d) = %d, expected %d",
				tt.tasks, tt.n, result, tt.expected)
		})
	}
}
