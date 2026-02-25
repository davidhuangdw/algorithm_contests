package main

import (
	"testing"
)

func TestBuildCartesianTree(t *testing.T) {
	// Test case 1: Simple increasing sequence
	// Input: [0, 1, 2, 3, 4] (1-indexed)
	// Expected: All left children should be 0, right children should form a chain
	w1 := []int{0, 1, 2, 3, 4}
	chd1 := buildCartesianTree(w1)

	expectedL1 := []int{0, 0, 0, 0, 0}
	expectedR1 := []int{0, 2, 3, 4, 0}

	for i := 1; i < len(w1); i++ {
		if chd1[i][0] != expectedL1[i] {
			t.Errorf("Test case 1 - left child of node %d: expected %d, got %d", i, expectedL1[i], chd1[i][0])
		}
		if chd1[i][1] != expectedR1[i] {
			t.Errorf("Test case 1 - right child of node %d: expected %d, got %d", i, expectedR1[i], chd1[i][1])
		}
	}

	// Test case 2: Simple decreasing sequence
	// Input: [0, 4, 3, 2, 1] (1-indexed)
	// Expected: All right children should be 0, left children should form a chain
	w2 := []int{0, 4, 3, 2, 1}
	chd2 := buildCartesianTree(w2)

	expectedL2 := []int{0, 0, 1, 2, 3}
	expectedR2 := []int{0, 0, 0, 0, 0}

	for i := 1; i < len(w2); i++ {
		if chd2[i][0] != expectedL2[i] {
			t.Errorf("Test case 2 - left child of node %d: expected %d, got %d", i, expectedL2[i], chd2[i][0])
		}
		if chd2[i][1] != expectedR2[i] {
			t.Errorf("Test case 2 - right child of node %d: expected %d, got %d", i, expectedR2[i], chd2[i][1])
		}
	}

	// Test case 3: Mixed sequence
	// Input: [0, 3, 1, 4, 2] (1-indexed)
	// Expected: Verify heap property (parent < children)
	w3 := []int{0, 3, 1, 4, 2}
	chd3 := buildCartesianTree(w3)

	// Verify heap property
	for i := 1; i < len(w3); i++ {
		if chd3[i][0] != 0 && w3[i] >= w3[chd3[i][0]] {
			t.Errorf("Test case 3 - Heap property violated: node %d (value %d) >= left child %d (value %d)",
				i, w3[i], chd3[i][0], w3[chd3[i][0]])
		}
		if chd3[i][1] != 0 && w3[i] >= w3[chd3[i][1]] {
			t.Errorf("Test case 3 - Heap property violated: node %d (value %d) >= right child %d (value %d)",
				i, w3[i], chd3[i][1], w3[chd3[i][1]])
		}
	}

	// Test case 4: Single element
	// Input: [0, 5] (1-indexed)
	// Expected: Both children should be 0
	w4 := []int{0, 5}
	chd4 := buildCartesianTree(w4)

	if chd4[1][0] != 0 || chd4[1][1] != 0 {
		t.Errorf("Test case 4 - Single element tree should have no children, got left=%d, right=%d", chd4[1][0], chd4[1][1])
	}

	// Test case 5: All elements equal
	// Input: [0, 2, 2, 2] (1-indexed)
	// Expected: Verify heap property (should still hold even with equal values)
	w5 := []int{0, 2, 2, 2}
	chd5 := buildCartesianTree(w5)

	// Verify heap property
	for i := 1; i < len(w5); i++ {
		if chd5[i][0] != 0 && w5[i] > w5[chd5[i][0]] {
			t.Errorf("Test case 5 - Heap property violated: node %d (value %d) > left child %d (value %d)",
				i, w5[i], chd5[i][0], w5[chd5[i][0]])
		}
		if chd5[i][1] != 0 && w5[i] > w5[chd5[i][1]] {
			t.Errorf("Test case 5 - Heap property violated: node %d (value %d) > right child %d (value %d)",
				i, w5[i], chd5[i][1], w5[chd5[i][1]])
		}
	}
}
