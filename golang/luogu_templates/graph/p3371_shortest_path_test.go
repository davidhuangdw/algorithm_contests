package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSPFA(t *testing.T) {
	// Test case 1: Simple directed graph with positive weights
	// 1 -> 2 (2)
	// 1 -> 3 (4)
	// 2 -> 3 (1)
	// 2 -> 4 (7)
	// 3 -> 4 (3)
	adj1 := make([][][2]int, 5) // nodes 1-4
	adj1[1] = [][2]int{{2, 2}, {3, 4}}
	adj1[2] = [][2]int{{3, 1}, {4, 7}}
	adj1[3] = [][2]int{{4, 3}}
	adj1[4] = [][2]int{}

	dis1 := spfa(4, 1, adj1)
	expected1 := []int{0, 0, 2, 3, 6} // indexes 0-4, node 0 unused
	assert.Equal(t, expected1, dis1, "Test case 1 failed")

	// Test case 2: Graph with all positive weights
	// 1 -> 2 (3)
	// 2 -> 3 (2)
	// 3 -> 2 (2)
	// 1 -> 4 (1)
	// 4 -> 5 (2)
	adj2 := make([][][2]int, 6) // nodes 1-5
	adj2[1] = [][2]int{{2, 3}, {4, 1}}
	adj2[2] = [][2]int{{3, 2}}
	adj2[3] = [][2]int{{2, 2}}
	adj2[4] = [][2]int{{5, 2}}
	adj2[5] = [][2]int{}

	dis2 := spfa(5, 1, adj2)
	expected2 := []int{0, 0, 3, 5, 1, 3} // indexes 0-5
	assert.Equal(t, expected2, dis2, "Test case 2 failed")

	// Test case 3: Single node graph
	adj3 := make([][][2]int, 2) // node 1
	adj3[1] = [][2]int{}

	dis3 := spfa(1, 1, adj3)
	expected3 := []int{0, 0}
	assert.Equal(t, expected3, dis3, "Test case 3 failed")

	// Test case 4: Chain-like graph
	// 1 -> 2 (1)
	// 2 -> 3 (1)
	// 3 -> 4 (1)
	// 4 -> 5 (1)
	adj4 := make([][][2]int, 6) // nodes 1-5
	adj4[1] = [][2]int{{2, 1}}
	adj4[2] = [][2]int{{3, 1}}
	adj4[3] = [][2]int{{4, 1}}
	adj4[4] = [][2]int{{5, 1}}
	adj4[5] = [][2]int{}

	dis4 := spfa(5, 1, adj4)
	expected4 := []int{0, 0, 1, 2, 3, 4}
	assert.Equal(t, expected4, dis4, "Test case 4 failed")

	// Test case 5: Graph with multiple paths between same nodes
	// 1 -> 2 (10)
	// 1 -> 3 (3)
	// 3 -> 2 (1)
	// 2 -> 4 (2)
	// 3 -> 4 (8)
	// 4 -> 5 (2)
	// 2 -> 5 (7)
	adj5 := make([][][2]int, 6) // nodes 1-5
	adj5[1] = [][2]int{{2, 10}, {3, 3}}
	adj5[2] = [][2]int{{4, 2}, {5, 7}}
	adj5[3] = [][2]int{{2, 1}, {4, 8}}
	adj5[4] = [][2]int{{5, 2}}
	adj5[5] = [][2]int{}

	dis5 := spfa(5, 1, adj5)
	expected5 := []int{0, 0, 4, 3, 6, 8} // shortest paths: 1->3->2->4->5 (3+1+2+2=8)
	assert.Equal(t, expected5, dis5, "Test case 5 failed")

	// Test case 6: Graph with unreachable nodes
	// 1 -> 2 (1)
	// 2 -> 3 (1)
	// 4 -> 5 (1)
	adj6 := make([][][2]int, 6) // nodes 1-5
	adj6[1] = [][2]int{{2, 1}}
	adj6[2] = [][2]int{{3, 1}}
	adj6[4] = [][2]int{{5, 1}}

	dis6 := spfa(5, 1, adj6)
	maxInt := 1<<31 - 1
	expected6 := []int{0, 0, 1, 2, maxInt, maxInt}
	assert.Equal(t, expected6, dis6, "Test case 6 failed")
}

func TestDijkstra(t *testing.T) {
	// Test case 1: Simple directed graph with positive weights
	// Same as SPFA test case 1 for comparison
	// 1 -> 2 (2)
	// 1 -> 3 (4)
	// 2 -> 3 (1)
	// 2 -> 4 (7)
	// 3 -> 4 (3)
	adj1 := make([][][2]int, 5) // nodes 1-4
	adj1[1] = [][2]int{{2, 2}, {3, 4}}
	adj1[2] = [][2]int{{3, 1}, {4, 7}}
	adj1[3] = [][2]int{{4, 3}}
	adj1[4] = [][2]int{}

	dis1 := dijkstra(4, 1, adj1)
	expected1 := []int{0, 0, 2, 3, 6} // indexes 0-4, node 0 unused
	assert.Equal(t, expected1, dis1, "Dijkstra test case 1 failed")

	// Test case 2: All positive weights graph
	// Same as SPFA test case 2 for comparison
	// 1 -> 2 (3)
	// 2 -> 3 (2)
	// 3 -> 2 (2)
	// 1 -> 4 (1)
	// 4 -> 5 (2)
	adj2 := make([][][2]int, 6) // nodes 1-5
	adj2[1] = [][2]int{{2, 3}, {4, 1}}
	adj2[2] = [][2]int{{3, 2}}
	adj2[3] = [][2]int{{2, 2}}
	adj2[4] = [][2]int{{5, 2}}
	adj2[5] = [][2]int{}

	dis2 := dijkstra(5, 1, adj2)
	expected2 := []int{0, 0, 3, 5, 1, 3} // indexes 0-5
	assert.Equal(t, expected2, dis2, "Dijkstra test case 2 failed")

	// Test case 3: Single node graph
	adj3 := make([][][2]int, 2) // node 1
	adj3[1] = [][2]int{}

	dis3 := dijkstra(1, 1, adj3)
	expected3 := []int{0, 0}
	assert.Equal(t, expected3, dis3, "Dijkstra test case 3 failed")

	// Test case 4: Chain-like graph
	// 1 -> 2 (1)
	// 2 -> 3 (1)
	// 3 -> 4 (1)
	// 4 -> 5 (1)
	adj4 := make([][][2]int, 6) // nodes 1-5
	adj4[1] = [][2]int{{2, 1}}
	adj4[2] = [][2]int{{3, 1}}
	adj4[3] = [][2]int{{4, 1}}
	adj4[4] = [][2]int{{5, 1}}
	adj4[5] = [][2]int{}

	dis4 := dijkstra(5, 1, adj4)
	expected4 := []int{0, 0, 1, 2, 3, 4}
	assert.Equal(t, expected4, dis4, "Dijkstra test case 4 failed")

	// Test case 5: Graph with multiple paths between same nodes
	// 1 -> 2 (10)
	// 1 -> 3 (3)
	// 3 -> 2 (1)
	// 2 -> 4 (2)
	// 3 -> 4 (8)
	// 4 -> 5 (2)
	// 2 -> 5 (7)
	adj5 := make([][][2]int, 6) // nodes 1-5
	adj5[1] = [][2]int{{2, 10}, {3, 3}}
	adj5[2] = [][2]int{{4, 2}, {5, 7}}
	adj5[3] = [][2]int{{2, 1}, {4, 8}}
	adj5[4] = [][2]int{{5, 2}}
	adj5[5] = [][2]int{}

	dis5 := dijkstra(5, 1, adj5)
	expected5 := []int{0, 0, 4, 3, 6, 8} // shortest paths: 1->3->2->4->5 (3+1+2+2=8)
	assert.Equal(t, expected5, dis5, "Dijkstra test case 5 failed")

	// Test case 6: Graph with unreachable nodes
	// 1 -> 2 (1)
	// 2 -> 3 (1)
	// 4 -> 5 (1)
	adj6 := make([][][2]int, 6) // nodes 1-5
	adj6[1] = [][2]int{{2, 1}}
	adj6[2] = [][2]int{{3, 1}}
	adj6[4] = [][2]int{{5, 1}}

	dis6 := dijkstra(5, 1, adj6)
	maxInt := 1<<31 - 1
	expected6 := []int{0, 0, 1, 2, maxInt, maxInt}
	assert.Equal(t, expected6, dis6, "Dijkstra test case 6 failed")

	// Test case 7: Complete graph with varying weights
	// 1 -> 2 (2), 1 -> 3 (5), 1 -> 4 (3)
	// 2 -> 1 (1), 2 -> 3 (2), 2 -> 4 (4)
	// 3 -> 1 (4), 3 -> 2 (3), 3 -> 4 (1)
	// 4 -> 1 (2), 4 -> 2 (5), 4 -> 3 (3)
	adj7 := make([][][2]int, 5) // nodes 1-4
	adj7[1] = [][2]int{{2, 2}, {3, 5}, {4, 3}}
	adj7[2] = [][2]int{{1, 1}, {3, 2}, {4, 4}}
	adj7[3] = [][2]int{{1, 4}, {2, 3}, {4, 1}}
	adj7[4] = [][2]int{{1, 2}, {2, 5}, {3, 3}}

	dis7 := dijkstra(4, 1, adj7)
	expected7 := []int{0, 0, 2, 4, 3} // indexes 0-4
	assert.Equal(t, expected7, dis7, "Dijkstra test case 7 failed")
}