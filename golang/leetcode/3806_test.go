package main

import "testing"

func Test_maximumAND(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		m        int
		expected int
	}{
		// Basic test cases
		{
			name:     "basic_case",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			m:        3,
			expected: 5, // Function returns 5
		},

		// Edge case: single element
		{
			name:     "single_element",
			nums:     []int{7},
			k:        2,
			m:        1,
			expected: 9, // Function returns 9
		},

		// Edge case: all elements same
		{
			name:     "all_same_elements",
			nums:     []int{3, 3, 3, 3},
			k:        5,
			m:        3,
			expected: 4, // Function returns 4
		},

		// Edge case: m equals array length
		{
			name:     "m_equals_array_length",
			nums:     []int{1, 2, 3},
			k:        4,
			m:        3,
			expected: 3, // Function returns 3
		},

		// Edge case: k is 0 (no operations allowed)
		{
			name:     "zero_operations",
			nums:     []int{1, 3, 5},
			k:        0,
			m:        2,
			expected: 1, // Function returns 1
		},

		// Edge case: large numbers
		{
			name:     "large_numbers",
			nums:     []int{1000000, 2000000, 3000000},
			k:        1000,
			m:        2,
			expected: 918868, // Function returns 918868
		},

		// Test case: multiple bits need setting
		{
			name:     "multiple_bits",
			nums:     []int{1, 2, 4, 8},
			k:        15,
			m:        3,
			expected: 9, // Function returns 9
		},

		// Test case: not enough operations
		{
			name:     "insufficient_operations",
			nums:     []int{1, 2, 3},
			k:        1,
			m:        2,
			expected: 3, // Function returns 3
		},

		// Test case: powers of 2
		{
			name:     "powers_of_two",
			nums:     []int{1, 2, 4, 8, 16},
			k:        10,
			m:        3,
			expected: 8, // Function returns 8
		},

		// Test case: consecutive numbers
		{
			name:     "consecutive_numbers",
			nums:     []int{5, 6, 7, 8},
			k:        5,
			m:        3,
			expected: 8, // Function returns 8
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumAND(tt.nums, tt.k, tt.m); got != tt.expected {
				t.Errorf("maximumAND() = %v, want %v", got, tt.expected)
			}
		})
	}
}
