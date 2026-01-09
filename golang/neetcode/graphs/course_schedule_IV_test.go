package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfPrerequisite(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		prerequisites [][]int
		queries       [][]int
		expected      []bool
	}{
		{
			name:          "no courses",
			n:             0,
			prerequisites: [][]int{},
			queries:       [][]int{},
			expected:      []bool{},
		},
		{
			name:          "single course no prerequisites",
			n:             1,
			prerequisites: [][]int{},
			queries:       [][]int{{0, 0}},
			expected:      []bool{false},
		},
		{
			name:          "direct prerequisite",
			n:             2,
			prerequisites: [][]int{{0, 1}},
			queries:       [][]int{{0, 1}},
			expected:      []bool{true},
		},
		{
			name:          "reverse query not prerequisite",
			n:             2,
			prerequisites: [][]int{{0, 1}},
			queries:       [][]int{{1, 0}},
			expected:      []bool{false},
		},
		{
			name:          "transitive prerequisite",
			n:             3,
			prerequisites: [][]int{{0, 1}, {1, 2}},
			queries:       [][]int{{0, 2}},
			expected:      []bool{true},
		},
		{
			name:          "leetcode example 1",
			n:             2,
			prerequisites: [][]int{{1, 0}},
			queries:       [][]int{{0, 1}, {1, 0}},
			expected:      []bool{false, true},
		},
		{
			name:          "leetcode example 2",
			n:             2,
			prerequisites: [][]int{},
			queries:       [][]int{{1, 0}, {0, 1}},
			expected:      []bool{false, false},
		},
		{
			name:          "multiple queries with cycles",
			n:             3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			queries:       [][]int{{0, 1}, {1, 2}, {2, 0}, {0, 2}, {1, 0}, {2, 1}},
			expected:      []bool{true, true, true, true, true, true},
		},
		{
			name:          "disconnected graph",
			n:             4,
			prerequisites: [][]int{{0, 1}, {2, 3}},
			queries:       [][]int{{0, 3}, {2, 1}, {0, 1}, {2, 3}},
			expected:      []bool{false, false, true, true},
		},
		{
			name:          "complex prerequisite chain",
			n:             5,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			queries:       [][]int{{0, 4}, {1, 3}, {2, 4}, {0, 2}, {3, 0}},
			expected:      []bool{true, true, true, true, false},
		},
		{
			name:          "self-loop not considered",
			n:             1,
			prerequisites: [][]int{{0, 0}},
			queries:       [][]int{{0, 0}},
			expected:      []bool{false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkIfPrerequisite(tt.n, tt.prerequisites, tt.queries)
			assert.Equal(t, tt.expected, result)
		})
	}
}