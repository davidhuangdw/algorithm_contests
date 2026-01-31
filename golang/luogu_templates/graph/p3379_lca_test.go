package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryLCA(t *testing.T) {
	// Test case 1: Simple tree with root 1
	// Structure: 1
	//          / | \
	//         2  3  4
	//        / \    |
	//       5   6   7
	adj1 := make([][]int, 8) // nodes 1-7
	adj1[1] = []int{2, 3, 4}
	adj1[2] = []int{1, 5, 6}
	adj1[3] = []int{1}
	adj1[4] = []int{1, 7}
	adj1[5] = []int{2}
	adj1[6] = []int{2}
	adj1[7] = []int{4}

	queries1 := [][2]int{
		{5, 6}, // LCA 2
		{5, 3}, // LCA 1
		{7, 3}, // LCA 1
		{4, 7}, // LCA 4
		{1, 5}, // LCA 1
		{2, 5}, // LCA 2
		{6, 7}, // LCA 1
	}

	ans1 := queryLCA(7, 1, adj1, queries1)
	expected1 := []int{2, 1, 1, 4, 1, 2, 1}
	assert.Equal(t, expected1, ans1, "Test case 1 failed")

	// Test case 2: Chain tree
	// Structure: 1 -> 2 -> 3 -> 4 -> 5
	adj2 := make([][]int, 6) // nodes 1-5
	adj2[1] = []int{2}
	adj2[2] = []int{1, 3}
	adj2[3] = []int{2, 4}
	adj2[4] = []int{3, 5}
	adj2[5] = []int{4}

	queries2 := [][2]int{
		{1, 5}, // LCA 1
		{2, 4}, // LCA 2
		{3, 5}, // LCA 3
		{4, 5}, // LCA 4
		{2, 2}, // LCA 2
	}

	ans2 := queryLCA(5, 1, adj2, queries2)
	expected2 := []int{1, 2, 3, 4, 2}
	assert.Equal(t, expected2, ans2, "Test case 2 failed")

	// Test case 3: Root with two branches
	// Structure: 1
	//          / \
	//         2   3
	//        / \   \
	//       4   5   6
	//      /     \   \
	//     7       8   9
	adj3 := make([][]int, 10) // nodes 1-9
	adj3[1] = []int{2, 3}
	adj3[2] = []int{1, 4, 5}
	adj3[3] = []int{1, 6}
	adj3[4] = []int{2, 7}
	adj3[5] = []int{2, 8}
	adj3[6] = []int{3, 9}
	adj3[7] = []int{4}
	adj3[8] = []int{5}
	adj3[9] = []int{6}

	queries3 := [][2]int{
		{7, 8},   // LCA 2
		{7, 9},   // LCA 1
		{4, 5},   // LCA 2
		{6, 9},   // LCA 6
		{1, 8},   // LCA 1
		{3, 5},   // LCA 1
		{7, 2},   // LCA 2
	}

	ans3 := queryLCA(9, 1, adj3, queries3)
	expected3 := []int{2, 1, 2, 6, 1, 1, 2}
	assert.Equal(t, expected3, ans3, "Test case 3 failed")

	// Test case 4: Single node tree
	adj4 := make([][]int, 2) // node 1
	adj4[1] = []int{}

	queries4 := [][2]int{
		{1, 1}, // LCA 1
	}

	ans4 := queryLCA(1, 1, adj4, queries4)
	expected4 := []int{1}
	assert.Equal(t, expected4, ans4, "Test case 4 failed")

	// Test case 5: Tree with different root
	// Same as test case 1 but rooted at 2 instead of 1
	adj5 := adj1 // same structure as test case 1
	queries5 := [][2]int{
		{5, 6}, // LCA 2
		{5, 1}, // LCA 2
		{7, 3}, // LCA 1
		{4, 7}, // LCA 4
		{2, 5}, // LCA 2
	}

	ans5 := queryLCA(7, 2, adj5, queries5)
	expected5 := []int{2, 2, 1, 4, 2}
	assert.Equal(t, expected5, ans5, "Test case 5 failed")
}