package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegTree(t *testing.T) {
	// Test case 1: Basic functionality
	initNums := []int{1, 2, 3, 4, 5}
	tr := NewSegTree(5, initNums)

	// Test initial sum
	assert.Equal(t, 15, tr.Sum(0, 4)) // Sum of [1,2,3,4,5]
	assert.Equal(t, 3, tr.Sum(0, 1))  // Sum of [1,2]
	assert.Equal(t, 7, tr.Sum(2, 3))  // Sum of [3,4]

	// Test Add operation
	tr.Add(1, 3, 2) // Add 2 to elements at indices 1, 2, 3
	// Now array should be [1, 4, 5, 6, 5]
	assert.Equal(t, 21, tr.Sum(0, 4)) // Sum of [1,4,5,6,5]
	assert.Equal(t, 1, tr.Sum(0, 0))  // Still 1
	assert.Equal(t, 4, tr.Sum(1, 1))  // Now 4
	assert.Equal(t, 11, tr.Sum(2, 3)) // Now [5,6] = 11
	assert.Equal(t, 15, tr.Sum(1, 3)) // Now [4,5,6] = 15

	// Test Add to single element
	tr.Add(4, 4, -3) // Add -3 to index 4
	// Now array should be [1, 4, 5, 6, 2]
	assert.Equal(t, 18, tr.Sum(0, 4)) // Sum of [1,4,5,6,2]
	assert.Equal(t, 2, tr.Sum(4, 4))  // Now 2

	// Test Add to entire range
	tr.Add(0, 4, 1) // Add 1 to all elements
	// Now array should be [2, 5, 6, 7, 3]
	assert.Equal(t, 23, tr.Sum(0, 4)) // Sum of [2,5,6,7,3]
	assert.Equal(t, 13, tr.Sum(0, 2)) // Sum of [2,5,6]

	// Test case 2: Empty tree
	tr2 := NewSegTree(0, []int{})
	assert.Equal(t, 0, tr2.Sum(0, -1))

	// Test case 3: Single element
	tr3 := NewSegTree(1, []int{10})
	assert.Equal(t, 10, tr3.Sum(0, 0))
	tr3.Add(0, 0, 5)
	assert.Equal(t, 15, tr3.Sum(0, 0))
}
