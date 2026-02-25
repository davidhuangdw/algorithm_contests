package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP5431(t *testing.T) {
	// Test with simple arrays
	assert.Equal(t, 1, p5431([]int{1}, 1, 2, 1), "Single element array should return 1")
	assert.Equal(t, 0, p5431([]int{1, 1}, 2, 3, 2), "Array [1,1] with p=3, k=2 should return 0")
	
	// Test with small arrays and prime modulus
	assert.Equal(t, 2, p5431([]int{1, 2, 3}, 3, 7, 2), "Array [1,2,3] with p=7, k=2 should return 2")
	assert.Equal(t, 3, p5431([]int{2, 3, 4}, 3, 11, 3), "Array [2,3,4] with p=11, k=3 should return 3")
	
	// Test with k=1 (simpler case)
	assert.Equal(t, 1, p5431([]int{1, 2, 3, 4}, 4, 13, 1), "Array [1,2,3,4] with p=13, k=1 should return 1")
	
	// Test with larger arrays
	assert.Equal(t, 533352, p5431([]int{1, 2, 3, 4, 5}, 5, 1000003, 2), "Array [1,2,3,4,5] with p=1000003, k=2 should return 533352")
	
	// Test with specific known values
	assert.Equal(t, 0, p5431([]int{1, 2, 3}, 3, 5, 2), "Array [1,2,3] with p=5, k=2 should return 0")
	assert.Equal(t, 12, p5431([]int{2, 4, 6, 8}, 4, 17, 3), "Array [2,4,6,8] with p=17, k=3 should return 12")
}