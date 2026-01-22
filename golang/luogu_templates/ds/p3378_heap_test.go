package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapPushTopPop(t *testing.T) {
	// Test basic functionality with integers
	h := NewHeap[int]()

	// Test pushing elements and checking top
	assert.Equal(t, 0, h.Len(), "Heap should be empty initially")
	h.Push(3)
	assert.Equal(t, 1, h.Len(), "Heap should have 1 element after push")
	assert.Equal(t, 3, h.Top(), "Top should be 3")

	h.Push(1)
	assert.Equal(t, 1, h.Top(), "Top should be 1 after pushing smaller element")

	h.Push(5)
	assert.Equal(t, 1, h.Top(), "Top should remain 1 after pushing larger element")

	// Test popping elements
	assert.Equal(t, 1, h.Pop(), "Pop should return smallest element 1")
	assert.Equal(t, 2, h.Len(), "Heap should have 2 elements after pop")
	assert.Equal(t, 3, h.Top(), "Top should now be 3")

	assert.Equal(t, 3, h.Pop(), "Pop should return next smallest element 3")
	assert.Equal(t, 1, h.Len(), "Heap should have 1 element after pop")
	assert.Equal(t, 5, h.Top(), "Top should now be 5")

	assert.Equal(t, 5, h.Pop(), "Pop should return last element 5")
	assert.Equal(t, 0, h.Len(), "Heap should be empty after popping all elements")
}

func TestHeapEdgeCases(t *testing.T) {
	// Test with single element
	h := NewHeap[int]()
	h.Push(42)
	assert.Equal(t, 42, h.Top(), "Top should be 42")
	assert.Equal(t, 42, h.Pop(), "Pop should return 42")
	assert.Equal(t, 0, h.Len(), "Heap should be empty")

	// Test with multiple pushes and pops
	h = NewHeap[int]()
	numbers := []int{7, 2, 9, 1, 5, 3}
	expectedOrder := []int{1, 2, 3, 5, 7, 9}

	for _, num := range numbers {
		h.Push(num)
	}

	for _, expected := range expectedOrder {
		assert.Equal(t, expected, h.Pop(), "Pop should return elements in ascending order")
	}

	assert.Equal(t, 0, h.Len(), "Heap should be empty after popping all elements")

	// Test with strings (since heap is generic)
	hStr := NewHeap[string]()
	strings := []string{"banana", "apple", "cherry", "date"}
	expectedStrOrder := []string{"apple", "banana", "cherry", "date"}

	for _, s := range strings {
		hStr.Push(s)
	}

	for _, expected := range expectedStrOrder {
		assert.Equal(t, expected, hStr.Pop(), "String heap should return elements in alphabetical order")
	}
}
