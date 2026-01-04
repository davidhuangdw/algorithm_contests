package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreqStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []int
	}{
		{
			name:       "LeetCode Example 1",
			operations: []string{"push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop"},
			values:     []int{5, 7, 5, 7, 4, 5, 0, 0, 0, 0},
			expected:   []int{0, 0, 0, 0, 0, 0, 5, 7, 5, 4},
		},
		{
			name:       "Single element",
			operations: []string{"push", "pop"},
			values:     []int{10, 0},
			expected:   []int{0, 10},
		},
		{
			name:       "Multiple same elements",
			operations: []string{"push", "push", "push", "pop", "pop", "pop"},
			values:     []int{5, 5, 5, 0, 0, 0},
			expected:   []int{0, 0, 0, 5, 5, 5},
		},
		{
			name:       "Tie breaking by recency",
			operations: []string{"push", "push", "push", "push", "pop", "pop", "pop", "pop"},
			values:     []int{1, 2, 1, 2, 0, 0, 0, 0},
			expected:   []int{0, 0, 0, 0, 2, 1, 2, 1},
		},
		{
			name:       "Complex frequency pattern",
			operations: []string{"push", "push", "push", "push", "push", "pop", "push", "pop", "pop", "pop"},
			values:     []int{1, 2, 3, 2, 1, 0, 3, 0, 0, 0},
			expected:   []int{0, 0, 0, 0, 0, 1, 0, 3, 2, 3},
		},
		{
			name:       "Empty stack pop",
			operations: []string{"pop"},
			values:     []int{0},
			expected:   []int{0},
		},
		{
			name:       "Multiple frequencies with same count",
			operations: []string{"push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop", "pop", "pop"},
			values:     []int{1, 2, 3, 1, 2, 3, 0, 0, 0, 0, 0, 0},
			expected:   []int{0, 0, 0, 0, 0, 0, 3, 2, 1, 3, 2, 1},
		},
		{
			name:       "Large numbers",
			operations: []string{"push", "push", "push", "pop", "pop", "pop"},
			values:     []int{1000, 2000, 1000, 0, 0, 0},
			expected:   []int{0, 0, 0, 1000, 2000, 1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := ConstructorFS()
			results := make([]int, 0, len(tt.operations))

			for i, op := range tt.operations {
				switch op {
				case "push":
					stack.Push(tt.values[i])
					results = append(results, 0) // Push returns nothing
				case "pop":
					result := stack.Pop()
					results = append(results, result)
				}
			}

			assert.Equal(t, tt.expected, results)
		})
	}
}
