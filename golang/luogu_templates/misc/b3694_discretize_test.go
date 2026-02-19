package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscretize(t *testing.T) {
	// Test with a simple array
	nums := []int{3, 1, 4, 1, 5}
	rank := discretize(nums)
	expected := map[int]int{1: 1, 3: 2, 4: 3, 5: 4}
	assert.Equal(t, expected, rank, "Simple array should be discretized correctly")

	// Test with already sorted array
	nums = []int{1, 2, 3, 4, 5}
	rank = discretize(nums)
	expected = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	assert.Equal(t, expected, rank, "Already sorted array should be discretized correctly")

	// Test with all same elements
	nums = []int{7, 7, 7, 7}
	rank = discretize(nums)
	expected = map[int]int{7: 1}
	assert.Equal(t, expected, rank, "Array with all same elements should be discretized correctly")

	// Test with negative numbers
	nums = []int{-3, -1, -2, -1, -4}
	rank = discretize(nums)
	expected = map[int]int{-4: 1, -3: 2, -2: 3, -1: 4}
	assert.Equal(t, expected, rank, "Array with negative numbers should be discretized correctly")

	// Test with mixed positive and negative numbers
	nums = []int{0, -2, 3, -2, 1}
	rank = discretize(nums)
	expected = map[int]int{-2: 1, 0: 2, 1: 3, 3: 4}
	assert.Equal(t, expected, rank, "Array with mixed numbers should be discretized correctly")

	// Test with single element
	nums = []int{42}
	rank = discretize(nums)
	expected = map[int]int{42: 1}
	assert.Equal(t, expected, rank, "Single element array should be discretized correctly")
}