package main

import "testing"

func TestLengthAfterTransformations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        int
		nums     []int
		expected int
	}{
		{
			name: "Example 1",
			s:    "abcyy",
			t:    2,
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			expected: 7,
		},
		{
			name: "Example 2",
			s:    "azbk",
			t:    1,
			nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			expected: 8,
		},
		{
			name: "Single char, t=0",
			s:    "a",
			t:    0,
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name: "All same, each produces 2",
			s:    "a",
			t:    1,
			nums: []int{3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthAfterTransformations(tt.s, tt.t, tt.nums)
			if result != tt.expected {
				t.Errorf("lengthAfterTransformations(%q, %v, %v) = %v; expected %v",
					tt.s, tt.t, tt.nums, result, tt.expected)
			}
		})
	}
}
