package main

import "testing"

func TestBalancedString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Already balanced",
			s:        "QWER",
			expected: 0,
		},
		{
			name:     "One extra Q",
			s:        "QQWE",
			expected: 1,
		},
		{
			name:     "Two extra Qs",
			s:        "QQQW",
			expected: 2,
		},
		{
			name:     "All Qs",
			s:        "QQQQ",
			expected: 3,
		},
		{
			name:     "Longer string with imbalance at the start",
			s:        "QQQQWER",
			expected: 3,
		},
		{
			name:     "Complex case from description",
			s:        "WWEQERQWQWWRWWERQWEQ",
			expected: 4, // This is a guess based on problem analysis, let's see if it holds
		},
		{
			name:     "Another complex case",
			s:        "WQWRQQQW",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedString(tt.s); got != tt.expected {
				t.Errorf("balancedString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
