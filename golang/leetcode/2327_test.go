package main

import "testing"

func TestPeopleAwareOfSecret(t *testing.T) {
	tests := []struct {
		n        int
		delay    int
		forget   int
		expected int
	}{
		{
			n:        6,
			delay:    2,
			forget:   4,
			expected: 5,
		},
		{
			n:        4,
			delay:    1,
			forget:   3,
			expected: 6,
		},
		{
			n:        1,
			delay:    1,
			forget:   2,
			expected: 1,
		},
		{
			n:        10,
			delay:    3,
			forget:   5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		if got := peopleAwareOfSecret(tt.n, tt.delay, tt.forget); got != tt.expected {
			t.Errorf("peopleAwareOfSecret(%d, %d, %d) = %v, want %v", tt.n, tt.delay, tt.forget, got, tt.expected)
		}
	}
}
