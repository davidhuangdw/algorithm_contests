package main

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		s         string
		t         string
		flipCost  int
		swapCost  int
		crossCost int
		expected  int64
	}{
		{
			s:         "01",
			t:         "01",
			flipCost:  1,
			swapCost:  10,
			crossCost: 100,
			expected:  0,
		},
		{
			s:         "0",
			t:         "1",
			flipCost:  1,
			swapCost:  10,
			crossCost: 100,
			expected:  1,
		},
		{
			s:         "01",
			t:         "10",
			flipCost:  10,
			swapCost:  1,
			crossCost: 100,
			expected:  1,
		},
		{
			s:         "00",
			t:         "11",
			flipCost:  10,
			swapCost:  1,
			crossCost: 1,
			expected:  2,
		},
		{
			s:         "0011",
			t:         "1100",
			flipCost:  1,
			swapCost:  10,
			crossCost: 2,
			expected:  4,
		},
	}

	for _, tt := range tests {
		if got := minimumCost(tt.s, tt.t, tt.flipCost, tt.swapCost, tt.crossCost); got != tt.expected {
			t.Errorf("minimumCost(%q, %q, %d, %d, %d) = %d, want %d",
				tt.s, tt.t, tt.flipCost, tt.swapCost, tt.crossCost, got, tt.expected)
		}
	}
}
