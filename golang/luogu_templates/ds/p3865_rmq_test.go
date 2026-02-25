package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRMQByST(t *testing.T) {
	// Test with a simple array
	nums := []int{1, 3, 2, 7, 5, 4, 6}
	queries := [][2]int{{0, 0}, {3, 3}, {0, 6}, {2, 4}, {4, 6}, {0, 2}}
	expected := []int{1, 7, 7, 7, 6, 3}

	results := rmqByST(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for simple array")

	// Test with array of size 1
	nums = []int{5}
	queries = [][2]int{{0, 0}}
	expected = []int{5}

	results = rmqByST(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for single-element array")

	// Test with array in descending order
	nums = []int{9, 7, 5, 3, 1}
	queries = [][2]int{{0, 4}, {2, 4}, {3, 3}}
	expected = []int{9, 5, 3}

	results = rmqByST(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for descending array")

	// Test with array in ascending order
	nums = []int{1, 3, 5, 7, 9}
	queries = [][2]int{{0, 4}, {2, 3}, {2, 2}}
	expected = []int{9, 7, 5}

	results = rmqByST(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for ascending array")

	// Test with duplicate values
	nums = []int{2, 5, 5, 2, 8, 8, 3}
	queries = [][2]int{{0, 6}, {1, 3}, {4, 5}}
	expected = []int{8, 5, 8}

	results = rmqByST(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for array with duplicates")
}

func TestRMQByCartesianLCA(t *testing.T) {
	// Test with a simple array
	nums := []int{1, 3, 2, 7, 5, 4, 6}
	queries := [][2]int{{0, 0}, {3, 3}, {0, 6}, {2, 4}, {4, 6}, {0, 2}}
	expected := []int{1, 7, 7, 7, 6, 3}

	results := rmqByCartesianLCA(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for simple array")

	// Test with array of size 1
	nums = []int{5}
	queries = [][2]int{{0, 0}}
	expected = []int{5}

	results = rmqByCartesianLCA(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for single-element array")

	// Test with array in descending order
	nums = []int{9, 7, 5, 3, 1}
	queries = [][2]int{{0, 4}, {2, 4}, {3, 3}}
	expected = []int{9, 5, 3}

	results = rmqByCartesianLCA(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for descending array")

	// Test with array in ascending order
	nums = []int{1, 3, 5, 7, 9}
	queries = [][2]int{{0, 4}, {2, 3}, {2, 2}}
	expected = []int{9, 7, 5}

	results = rmqByCartesianLCA(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for ascending array")

	// Test with duplicate values
	nums = []int{2, 5, 5, 2, 8, 8, 3}
	queries = [][2]int{{0, 6}, {1, 3}, {4, 5}}
	expected = []int{8, 5, 8}

	results = rmqByCartesianLCA(nums, queries)
	assert.Equal(t, expected, results, "Results should match expected values for array with duplicates")
}

// Test that both implementations produce the same results
func TestRMQConsistency(t *testing.T) {
	// Test with various arrays to ensure both methods produce consistent results
	testCases := [][]int{
		{1, 3, 2, 7, 5, 4, 6},    // mixed array
		{5},                      // single element
		{9, 7, 5, 3, 1},          // descending
		{1, 3, 5, 7, 9},          // ascending
		{2, 5, 5, 2, 8, 8, 3},    // duplicates
		{3, 1, 4, 1, 5, 9, 2, 6}, // another mixed array
	}

	for _, nums := range testCases {
		// Generate various queries for each array
		queries := [][2]int{}
		queries = append(queries, [2]int{0, len(nums) - 1}) // full range
		queries = append(queries, [2]int{0, 0})             // first element
		queries = append(queries, [2]int{len(nums) - 1, len(nums) - 1}) // last element
		
		// Add some random range queries
		if len(nums) > 2 {
			queries = append(queries, [2]int{0, 2})
			queries = append(queries, [2]int{len(nums) - 3, len(nums) - 1})
		}

		// Get results from both implementations
		resultsST := rmqByST(nums, queries)
		resultsLCA := rmqByCartesianLCA(nums, queries)

		// Verify results are identical
		assert.Equal(t, resultsST, resultsLCA, "Both RMQ implementations should produce identical results for array %v", nums)
	}
}