package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildMatrix(t *testing.T) {
	tests := []struct {
		name          string
		k             int
		rowConditions [][]int
		colConditions [][]int
		expected      [][]int
		shouldError   bool
		description   string
	}{
		{
			name:          "LeetCode Example 1",
			k:             3,
			rowConditions: [][]int{{1, 2}, {3, 2}},
			colConditions: [][]int{{2, 1}, {3, 2}},
			expected:      [][]int{{0, 0, 1}, {3, 0, 0}, {0, 2, 0}},
			shouldError:   false,
			description:   "Standard case with both row and column conditions",
		},
		{
			name:          "LeetCode Example 2",
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}},
			colConditions: [][]int{{2, 1}},
			expected:      nil,
			shouldError:   true,
			description:   "Row conditions contain a cycle",
		},
		{
			name:          "Single element",
			k:             1,
			rowConditions: [][]int{},
			colConditions: [][]int{},
			expected:      [][]int{{1}},
			shouldError:   false,
			description:   "Only one element, no conditions",
		},
		{
			name:          "No conditions",
			k:             2,
			rowConditions: [][]int{},
			colConditions: [][]int{},
			expected:      [][]int{{2, 1}, {0, 0}},
			shouldError:   false,
			description:   "No conditions, elements can be placed in any order",
		},
		{
			name:          "Column cycle",
			k:             2,
			rowConditions: [][]int{{1, 2}},
			colConditions: [][]int{{1, 2}, {2, 1}},
			expected:      nil,
			shouldError:   true,
			description:   "Column conditions contain a cycle",
		},
		{
			name:          "Only row conditions",
			k:             3,
			rowConditions: [][]int{{1, 2}, {1, 3}},
			colConditions: [][]int{},
			expected:      [][]int{{2, 3, 1}, {0, 0, 0}, {0, 0, 0}},
			shouldError:   false,
			description:   "Only row conditions, no column conditions",
		},
		{
			name:          "Only column conditions",
			k:             3,
			rowConditions: [][]int{},
			colConditions: [][]int{{1, 2}, {1, 3}},
			expected:      [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}},
			shouldError:   false,
			description:   "Only column conditions, no row conditions",
		},
		{
			name:          "Complex dependencies",
			k:             4,
			rowConditions: [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}},
			colConditions: [][]int{{1, 2}, {3, 2}, {3, 4}},
			expected:      [][]int{{0, 0, 0, 1}, {0, 0, 3, 0}, {4, 0, 0, 0}, {0, 2, 0, 0}},
			shouldError:   false,
			description:   "Complex dependencies in both rows and columns",
		},
		{
			name:          "No overlapping conditions",
			k:             4,
			rowConditions: [][]int{{1, 2}, {3, 4}},
			colConditions: [][]int{{1, 3}, {2, 4}},
			expected:      [][]int{{0, 0, 1, 0}, {0, 0, 0, 2}, {0, 3, 0, 0}, {4, 0, 0, 0}},
			shouldError:   false,
			description:   "Conditions form separate chains",
		},
		{
			name:          "Identity conditions",
			k:             2,
			rowConditions: [][]int{{1, 2}},
			colConditions: [][]int{{1, 2}},
			expected:      [][]int{{0, 1}, {2, 0}},
			shouldError:   false,
			description:   "Same conditions for rows and columns",
		},
		{
			name:          "All elements in same row",
			k:             3,
			rowConditions: [][]int{},
			colConditions: [][]int{{1, 2}, {2, 3}},
			expected:      [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}},
			shouldError:   false,
			description:   "No row conditions, only column order",
		},
		{
			name:          "All elements in same column",
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}},
			colConditions: [][]int{},
			expected:      [][]int{{0, 0, 1}, {0, 0, 2}, {0, 0, 3}},
			shouldError:   false,
			description:   "No column conditions, only row order",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildMatrix(tt.k, tt.rowConditions, tt.colConditions)

			if tt.shouldError {
				assert.Nil(t, result, "Expected nil due to cycle in conditions")
			} else {
				// Check if the matrix has correct dimensions
				assert.Equal(t, tt.k, len(result), "Matrix should have %d rows, got %d", tt.k, len(result))
				for i, row := range result {
					assert.Equal(t, tt.k, len(row), "Row %d should have %d columns, got %d", i, tt.k, len(row))
				}

				// Check if all elements 1..k are present exactly once
				elements := make(map[int]bool)
				for i := 0; i < tt.k; i++ {
					for j := 0; j < tt.k; j++ {
						val := result[i][j]
						if val > 0 && val <= tt.k {
							elements[val] = true
						}
					}
				}
				assert.Equal(t, tt.k, len(elements), "Should contain exactly %d distinct elements from 1 to %d, got %d", tt.k, tt.k, len(elements))

				// Check row conditions are satisfied
				for _, cond := range tt.rowConditions {
					above, below := cond[0], cond[1]
					var aboveRow, belowRow int
					for i := 0; i < tt.k; i++ {
						for j := 0; j < tt.k; j++ {
							if result[i][j] == above {
								aboveRow = i
							}
							if result[i][j] == below {
								belowRow = i
							}
						}
					}
					assert.Less(t, aboveRow, belowRow, "Element %d should be above %d, but they're at rows %d and %d", above, below, aboveRow, belowRow)
				}

				// Check column conditions are satisfied
				for _, cond := range tt.colConditions {
					left, right := cond[0], cond[1]
					var leftCol, rightCol int
					for i := 0; i < tt.k; i++ {
						for j := 0; j < tt.k; j++ {
							if result[i][j] == left {
								leftCol = j
							}
							if result[i][j] == right {
								rightCol = j
							}
						}
					}
					assert.Less(t, leftCol, rightCol, "Element %d should be left of %d, but they're at columns %d and %d", left, right, leftCol, rightCol)
				}
			}
		})
	}
}
