package main

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected []int
	}{
		{
			name:     "Even number of nodes",
			head:     toList([]int{1, 2, 3, 4}),
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "Odd number of nodes",
			head:     toList([]int{1, 2, 3, 4, 5}),
			expected: []int{1, 5, 2, 4, 3},
		},
		{
			name:     "Single node",
			head:     toList([]int{1}),
			expected: []int{1},
		},
		{
			name:     "Two nodes",
			head:     toList([]int{1, 2}),
			expected: []int{1, 2},
		},
		{
			name:     "Empty list",
			head:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.head)
			if got := fromList(tt.head); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reorderList() = %v, want %v", got, tt.expected)
			}
		})
	}
}
