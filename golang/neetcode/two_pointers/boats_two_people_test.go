package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name     string
		people   []int
		limit    int
		expected int
	}{
		{
			name:     "leetcode example 1",
			people:   []int{1, 2},
			limit:    3,
			expected: 1, // both can fit in one boat
		},
		{
			name:     "leetcode example 2",
			people:   []int{3, 2, 2, 1},
			limit:    3,
			expected: 3, // [3], [2,1], [2]
		},
		{
			name:     "leetcode example 3",
			people:   []int{3, 5, 3, 4},
			limit:    5,
			expected: 4, // each person needs their own boat
		},
		{
			name:     "single person",
			people:   []int{5},
			limit:    10,
			expected: 1,
		},
		{
			name:     "empty array",
			people:   []int{},
			limit:    5,
			expected: 0,
		},
		{
			name:     "all people can pair",
			people:   []int{1, 1, 1, 1, 1, 1},
			limit:    2,
			expected: 3, // 3 boats with 2 people each
		},
		{
			name:     "no pairing possible",
			people:   []int{4, 4, 4, 4},
			limit:    7,
			expected: 4, // each needs their own boat
		},
		{
			name:     "mixed weights",
			people:   []int{1, 3, 4, 2, 5},
			limit:    6,
			expected: 3, // [5,1], [4,2], [3]
		},
		{
			name:     "large limit",
			people:   []int{10, 20, 30},
			limit:    100,
			expected: 2, // [30,20], [10]
		},
		{
			name:     "exact pairing",
			people:   []int{2, 3, 4, 1, 1, 2},
			limit:    5,
			expected: 3, // [4,1], [3,2], [2,1]
		},
		{
			name:     "heavy people",
			people:   []int{100, 200, 150, 50},
			limit:    200,
			expected: 3, // [200], [150,50], [100]
		},
		{
			name:     "duplicate weights",
			people:   []int{2, 2, 2, 2, 2},
			limit:    4,
			expected: 3, // [2,2], [2,2], [2]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numRescueBoats(tt.people, tt.limit)
			assert.Equal(t, tt.expected, result, "people=%v, limit=%d", tt.people, tt.limit)
		})
	}
}