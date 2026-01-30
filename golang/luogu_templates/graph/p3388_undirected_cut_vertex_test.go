package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCuts(t *testing.T) {
	// Test case 1: Single node
	adj1 := [][]int{nil, []int{}}
	cuts1 := getCuts(1, adj1)
	assert.Empty(t, cuts1)

	// Test case 2: Two nodes connected
	adj2 := [][]int{nil, []int{2}, []int{1}}
	cuts2 := getCuts(2, adj2)
	assert.Empty(t, cuts2)

	// Test case 3: Chain of three nodes (middle is cut vertex)
	adj3 := [][]int{nil, []int{2}, []int{1, 3}, []int{2}}
	cuts3 := getCuts(3, adj3)
	assert.Equal(t, map[int]bool{2: true}, cuts3)

	// Test case 4: Root with multiple children (root is cut vertex)
	adj4 := [][]int{nil, []int{2, 3, 4}, []int{1}, []int{1}, []int{1}}
	cuts4 := getCuts(4, adj4)
	assert.Equal(t, map[int]bool{1: true}, cuts4)

	// Test case 5: Root with single child (root is not cut vertex)
	adj5 := [][]int{nil, []int{2}, []int{1, 3, 4, 5}, []int{2}, []int{2}, []int{2}}
	cuts5 := getCuts(5, adj5)
	assert.Equal(t, map[int]bool{2: true}, cuts5)

	// Test case 6: Graph with cycle (2 is cut vertex)
	adj6 := [][]int{nil, []int{2, 4}, []int{1, 3, 5}, []int{2}, []int{1, 5}, []int{2, 4}}
	cuts6 := getCuts(5, adj6)
	assert.Equal(t, map[int]bool{2: true}, cuts6)

	// Test case 7: Multiple cut vertices (2, 3, 4)
	adj7 := [][]int{nil, []int{2}, []int{1, 3}, []int{2, 4}, []int{3, 5}, []int{4}}
	cuts7 := getCuts(5, adj7)
	assert.Equal(t, map[int]bool{2: true, 3: true, 4: true}, cuts7)

	// Test case 8: Tree structure with multiple cut vertices (1, 2, 3)
	adj8 := [][]int{nil, []int{2, 3}, []int{1, 4, 5, 6}, []int{1, 7}, []int{2}, []int{2}, []int{2}, []int{3}}
	cuts8 := getCuts(7, adj8)
	assert.Equal(t, map[int]bool{1: true, 2: true, 3: true}, cuts8)

	// Test case 9: Multiple edges between same nodes (should not affect result)
	adj9 := [][]int{nil, []int{2, 2}, []int{1, 1, 3}, []int{2}}
	cuts9 := getCuts(3, adj9)
	assert.Equal(t, map[int]bool{2: true}, cuts9)

	// Test case 10: Graph where node has back edge to grandparent (2 should be cut vertex)
	adj10 := [][]int{nil, []int{2, 4}, []int{1, 3, 5}, []int{2}, []int{1, 5}, []int{2, 4}}
	cuts10 := getCuts(5, adj10)
	assert.Equal(t, map[int]bool{2: true}, cuts10)
}
