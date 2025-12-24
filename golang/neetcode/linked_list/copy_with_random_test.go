package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a linked list with random pointers
func createRandomList(values []int, randomIndices []int) *Node {
	if len(values) == 0 {
		return nil
	}
	
	nodes := make([]*Node, len(values))
	
	// Create all nodes first
	for i, val := range values {
		nodes[i] = &Node{Val: val}
	}
	
	// Set next pointers
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	
	// Set random pointers
	for i, randomIdx := range randomIndices {
		if randomIdx >= 0 && randomIdx < len(nodes) {
			nodes[i].Random = nodes[randomIdx]
		}
	}
	
	return nodes[0]
}

// Helper function to compare two linked lists with random pointers
func compareRandomLists(l1 *Node, l2 *Node) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	
	// Create mapping from original to copy
	originalToCopy := make(map[*Node]*Node)
	p1, p2 := l1, l2
	
	// First pass: check values and create mapping
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		originalToCopy[p1] = p2
		p1 = p1.Next
		p2 = p2.Next
	}
	
	// Check if both lists have same length
	if p1 != nil || p2 != nil {
		return false
	}
	
	// Second pass: check random pointers
	p1, p2 = l1, l2
	for p1 != nil && p2 != nil {
		if p1.Random == nil {
			if p2.Random != nil {
				return false
			}
		} else {
			expectedRandom := originalToCopy[p1.Random]
			if p2.Random != expectedRandom {
				return false
			}
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	
	return true
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name          string
		values        []int
		randomIndices []int
	}{
		// Basic cases
		{"empty list", []int{}, []int{}},
		{"single element no random", []int{1}, []int{-1}},
		{"single element self random", []int{1}, []int{0}},
		
		// Classic cases
		{"two elements no random", []int{1, 2}, []int{-1, -1}},
		{"two elements random to each other", []int{1, 2}, []int{1, 0}},
		{"three elements chain random", []int{1, 2, 3}, []int{1, 2, 0}},
		{"three elements all to head", []int{1, 2, 3}, []int{0, 0, 0}},
		
		// Edge cases
		{"all same values", []int{1, 1, 1, 1}, []int{2, 3, 0, 1}},
		{"negative numbers", []int{-1, -2, -3}, []int{2, 0, 1}},
		{"mixed numbers", []int{-1, 0, 1}, []int{2, 1, 0}},
		
		// Complex cases
		{"large list sequential random", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 0}},
		{"large list reverse random", []int{1, 2, 3, 4, 5}, []int{4, 3, 2, 1, 0}},
		{"large list all to same node", []int{1, 2, 3, 4, 5}, []int{2, 2, 2, 2, 2}},
		{"some nil random pointers", []int{1, 2, 3, 4, 5}, []int{-1, 0, -1, 2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := createRandomList(tt.values, tt.randomIndices)
			copy := copyRandomList(original)
			
			// Verify the copy is correct
			assert.True(t, compareRandomLists(original, copy), "Copied list should match original structure")
			
			// Verify it's a deep copy (not same memory addresses)
			if original != nil && copy != nil {
				assert.NotSame(t, original, copy, "Copy should be a different object")
				
				// Check that no nodes are shared between original and copy
				originalNodes := make(map[*Node]bool)
				for p := original; p != nil; p = p.Next {
					originalNodes[p] = true
				}
				
				for p := copy; p != nil; p = p.Next {
					assert.False(t, originalNodes[p], "Copy should not share nodes with original")
				}
			}
		})
	}
}