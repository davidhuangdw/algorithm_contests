package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// pos indicates the 0-based index where the cycle starts (-1 for no cycle)
func createListWithCycle(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))
	head := &ListNode{Val: values[0]}
	nodes[0] = head
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		nodes[i] = current
	}

	// Create cycle if pos is valid
	if pos >= 0 && pos < len(nodes) {
		current.Next = nodes[pos]
	}

	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int
		expected bool
	}{
		// Basic cases
		{"empty list", []int{}, -1, false},
		{"single element no cycle", []int{1}, -1, false},
		{"single element self-cycle", []int{1}, 0, true},

		// Classic cases
		{"two elements no cycle", []int{1, 2}, -1, false},
		{"two elements cycle", []int{1, 2}, 0, true},
		{"three elements no cycle", []int{1, 2, 3}, -1, false},
		{"three elements cycle at start", []int{1, 2, 3}, 0, true},
		{"three elements cycle at middle", []int{1, 2, 3}, 1, true},
		{"three elements cycle at end", []int{1, 2, 3}, 2, true},

		// Edge cases
		{"four elements cycle at start", []int{1, 2, 3, 4}, 0, true},
		{"four elements cycle at middle", []int{1, 2, 3, 4}, 1, true},
		{"four elements cycle at second middle", []int{1, 2, 3, 4}, 2, true},
		{"four elements cycle at end", []int{1, 2, 3, 4}, 3, true},

		// Complex cases
		{"large list no cycle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, false},
		{"large list cycle at start", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, true},
		{"large list cycle at middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, true},
		{"large list cycle at end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, true},
		{"cycle with duplicate values", []int{1, 2, 2, 3, 4}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createListWithCycle(tt.values, tt.pos)
			result := hasCycle(list)
			assert.Equal(t, tt.expected, result, "Cycle detection should match expected")
		})
	}
}
