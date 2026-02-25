package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCantorRank(t *testing.T) {
	// Test with simple permutations
	assert.Equal(t, 1, cantorRank([]int{1}), "Single element permutation should have rank 1")
	assert.Equal(t, 1, cantorRank([]int{1, 2}), "Permutation [1,2] should have rank 1")
	assert.Equal(t, 2, cantorRank([]int{2, 1}), "Permutation [2,1] should have rank 2")
	
	// Test with 3-element permutations
	assert.Equal(t, 1, cantorRank([]int{1, 2, 3}), "Permutation [1,2,3] should have rank 1")
	assert.Equal(t, 2, cantorRank([]int{1, 3, 2}), "Permutation [1,3,2] should have rank 2")
	assert.Equal(t, 3, cantorRank([]int{2, 1, 3}), "Permutation [2,1,3] should have rank 3")
	assert.Equal(t, 4, cantorRank([]int{2, 3, 1}), "Permutation [2,3,1] should have rank 4")
	assert.Equal(t, 5, cantorRank([]int{3, 1, 2}), "Permutation [3,1,2] should have rank 5")
	assert.Equal(t, 6, cantorRank([]int{3, 2, 1}), "Permutation [3,2,1] should have rank 6")
	
	// Test with larger permutations
	assert.Equal(t, 1, cantorRank([]int{1, 2, 3, 4}), "Permutation [1,2,3,4] should have rank 1")
	assert.Equal(t, 24, cantorRank([]int{4, 3, 2, 1}), "Permutation [4,3,2,1] should have rank 24 (4!)")
	
	// Test with specific known values
	assert.Equal(t, 8, cantorRank([]int{2, 1, 4, 3}), "Permutation [2,1,4,3] should have rank 8")
	assert.Equal(t, 12, cantorRank([]int{1, 3, 5, 4, 2}), "Permutation [1,3,5,4,2] should have rank 12")
}