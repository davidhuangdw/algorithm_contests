package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartitionKSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{{
		name:     "Basic case - can partition into 3 subsets",
		nums:     []int{4, 3, 2, 3, 5, 2, 1},
		k:        4,
		expected: true,
	}, {
		name:     "Basic case",
		nums:     []int{2, 2, 2, 1, 1},
		k:        2,
		expected: true,
	}, {
		name:     "Cannot partition - sum not divisible by k",
		nums:     []int{1, 2, 3, 4},
		k:        3,
		expected: false,
	}, {
		name:     "Cannot partition - individual element too large",
		nums:     []int{10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6},
		k:        3,
		expected: true,
	}, {
		name:     "Edge case - k=1",
		nums:     []int{1, 2, 3, 4},
		k:        1,
		expected: true,
	}, {
		name:     "Edge case - k equals number of elements",
		nums:     []int{1, 1, 1, 1},
		k:        4,
		expected: true,
	}, {
		name:     "Edge case - k=2 with two elements",
		nums:     []int{5, 5},
		k:        2,
		expected: true,
	}, {
		name:     "Complex case - cannot partition",
		nums:     []int{2, 2, 2, 2, 3, 4, 5},
		k:        4,
		expected: false,
	}, {
		name:     "Complex case - cannot partition",
		nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		k:        3,
		expected: false,
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canPartitionKSubsets(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result, "canPartitionKSubsets(%v, %d) should be %v", tt.nums, tt.k, tt.expected)
		})
	}
}
