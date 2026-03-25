package main

import "testing"

func TestCountBalanced(t *testing.T) {
	tests := []struct {
		name     string
		low      int64
		high     int64
		expected int64
	}{
		{
			name:     "Example 1",
			low:      1,
			high:     100,
			expected: 9, // 11, 22, 33, 44, 55, 66, 77, 88, 99
		},
		{
			name:     "Example 2",
			low:      10,
			high:     11,
			expected: 1, // 11
		},
		{
			name:     "Range with no balanced numbers",
			low:      12,
			high:     21,
			expected: 0,
		},
		{
			name:     "Larger range",
			low:      1,
			high:     1000,
			expected: 54,
		},
		{
			name:     "Low greater than high",
			low:      4,
			high:     1000000000000000,
			expected: 32811494188974,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalanced(tt.low, tt.high); got != tt.expected {
				t.Errorf("countBalanced() = %v, want %v", got, tt.expected)
			}
		})
	}
}
