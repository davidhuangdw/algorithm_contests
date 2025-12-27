package backtrace

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{-1, 0, 1}, {-1, 1, 0}, {0, -1, 1},
				{0, 1, -1}, {1, -1, 0}, {1, 0, -1},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			expected: [][]int{
				{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
				{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
				{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
				{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permute(tt.nums)

			// Sort both result and expected for consistent comparison
			sortPermutations(result)
			sortPermutations(tt.expected)

			assert.Equal(t, tt.expected, result,
				"permute(%v) should equal %v", tt.nums, tt.expected)
		})
	}
}

// Helper function to sort permutations for consistent comparison
func sortPermutations(permutations [][]int) {
	// Sort each permutation
	for _, perm := range permutations {
		sort.Ints(perm)
	}

	// Sort the list of permutations
	sort.Slice(permutations, func(i, j int) bool {
		for k := 0; k < len(permutations[i]) && k < len(permutations[j]); k++ {
			if permutations[i][k] != permutations[j][k] {
				return permutations[i][k] < permutations[j][k]
			}
		}
		return len(permutations[i]) < len(permutations[j])
	})
}
