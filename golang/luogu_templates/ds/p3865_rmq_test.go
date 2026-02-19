package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRMQ(t *testing.T) {
	// Test with a simple array
	nums := []int{1, 3, 2, 7, 5, 4, 6}
	rmq := NewRMQ(nums)

	// Test single element queries
	assert.Equal(t, 1, rmq.query(0, 0), "Query [0,0] should return 1")
	assert.Equal(t, 7, rmq.query(3, 3), "Query [3,3] should return 7")

	// Test range queries
	assert.Equal(t, 7, rmq.query(0, 6), "Query [0,6] should return 7")
	assert.Equal(t, 7, rmq.query(2, 4), "Query [2,4] should return 7")
	assert.Equal(t, 6, rmq.query(4, 6), "Query [4,6] should return 6")
	assert.Equal(t, 3, rmq.query(0, 2), "Query [0,2] should return 3")

	// Test with array of size 1
	nums = []int{5}
	rmq = NewRMQ(nums)
	assert.Equal(t, 5, rmq.query(0, 0), "Query [0,0] should return 5")

	// Test with array in descending order
	nums = []int{9, 7, 5, 3, 1}
	rmq = NewRMQ(nums)
	assert.Equal(t, 9, rmq.query(0, 4), "Query [0,4] should return 9")
	assert.Equal(t, 5, rmq.query(2, 4), "Query [2,4] should return 5")
	assert.Equal(t, 3, rmq.query(3, 3), "Query [3,3] should return 3")

	// Test with array in ascending order
	nums = []int{1, 3, 5, 7, 9}
	rmq = NewRMQ(nums)
	assert.Equal(t, 9, rmq.query(0, 4), "Query [0,4] should return 9")
	assert.Equal(t, 7, rmq.query(2, 3), "Query [2,3] should return 7")
	assert.Equal(t, 5, rmq.query(2, 2), "Query [2,2] should return 5")

	// Test with duplicate values
	nums = []int{2, 5, 5, 2, 8, 8, 3}
	rmq = NewRMQ(nums)
	assert.Equal(t, 8, rmq.query(0, 6), "Query [0,6] should return 8")
	assert.Equal(t, 5, rmq.query(1, 3), "Query [1,3] should return 5")
	assert.Equal(t, 8, rmq.query(4, 5), "Query [4,5] should return 8")
}