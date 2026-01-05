package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to compare two quad trees
func compareQuadTrees(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.IsLeaf != b.IsLeaf || a.Val != b.Val {
		return false
	}
	if a.IsLeaf {
		return true // Leaf nodes don't have children
	}
	return compareQuadTrees(a.TopLeft, b.TopLeft) &&
		compareQuadTrees(a.TopRight, b.TopRight) &&
		compareQuadTrees(a.BottomLeft, b.BottomLeft) &&
		compareQuadTrees(a.BottomRight, b.BottomRight)
}

func TestConstruct(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected *Node
	}{
		// Basic cases
		{
			name: "1x1 grid with 0",
			grid: [][]int{{0}},
			expected: &Node{Val: false, IsLeaf: true},
		},
		{
			name: "1x1 grid with 1",
			grid: [][]int{{1}},
			expected: &Node{Val: true, IsLeaf: true},
		},

		// 2x2 grids
		{
			name: "2x2 uniform grid with 0",
			grid: [][]int{{0, 0}, {0, 0}},
			expected: &Node{Val: false, IsLeaf: true},
		},
		{
			name: "2x2 uniform grid with 1",
			grid: [][]int{{1, 1}, {1, 1}},
			expected: &Node{Val: true, IsLeaf: true},
		},
		{
			name: "2x2 mixed grid",
			grid: [][]int{{0, 1}, {1, 0}},
			expected: &Node{
				IsLeaf: false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},

		// 4x4 grids
		{
			name: "4x4 uniform grid with 1",
			grid: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: &Node{Val: true, IsLeaf: true},
		},
		{
			name: "4x4 grid with top-left quadrant different",
			grid: [][]int{
				{0, 0, 1, 1},
				{0, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: &Node{
				IsLeaf: false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: true, IsLeaf: true},
			},
		},

		// Edge cases
		{
			name: "empty grid",
			grid: [][]int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := construct(tt.grid)
			
			// Compare the resulting quad tree with expected
			assert.True(t, compareQuadTrees(result, tt.expected),
				"construct(%v) should produce the expected quad tree structure",
				tt.grid)
		})
	}
}