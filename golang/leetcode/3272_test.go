package main

import "testing"

func TestCountGoodIntegers(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected int64
	}{
		{
			name:     "n=1, k=1",
			n:        1,
			k:        1,
			expected: 9,
		},
		{
			name:     "n=2, k=11",
			n:        2,
			k:        11,
			expected: 9,
		},
		{
			name:     "n=2, k=3",
			n:        2,
			k:        3,
			expected: 3,
		},
		{
			name:     "n=3, k=101",
			n:        3,
			k:        101,
			expected: 18,
		},
		{
			name:     "n=4, k=1",
			n:        4,
			k:        1,
			expected: 252,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodIntegers(tt.n, tt.k); got != tt.expected {
				t.Errorf("countGoodIntegers() = %v, want %v", got, tt.expected)
			}
		})
	}
}
