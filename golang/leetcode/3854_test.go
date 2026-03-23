package main

import (
	"reflect"
	"testing"
)

func TestMakeParityAlternating(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1: Already alternating in one form",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{0, 4}, // 0 moves needed for O,E,O,E,O pattern.
		},
		{
			name:     "Example 2",
			nums:     []int{6, 0, 3, 2, 5},
			expected: []int{1, 5}, // 1 move needed for O,E,O,E,O pattern.
		},
		{
			name:     "Already alternating",
			nums:     []int{0, 1, 0, 1, 0},
			expected: []int{0, 1},
		},
		{
			name:     "Single element",
			nums:     []int{10},
			expected: []int{0, 0},
		},
		{
			name:     "All same parity",
			nums:     []int{2, 4, 6, 8},
			expected: []int{2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeParityAlternating(tt.nums); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("makeParityAlternating() = %v, want %v", got, tt.expected)
			}
		})
	}
}
