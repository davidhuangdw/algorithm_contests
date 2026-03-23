package main

import "testing"

func TestMaxPower(t *testing.T) {
	tests := []struct {
		name     string
		stations []int
		rad      int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			stations: []int{1, 2, 4, 5, 0},
			rad:      1,
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			stations: []int{4, 4, 4, 4},
			rad:      0,
			k:        3,
			expected: 4,
		},
		{
			name:     "Single station, large k",
			stations: []int{10},
			rad:      0,
			k:        10,
			expected: 20,
		},
		{
			name:     "Zero radius, needs distribution",
			stations: []int{1, 0, 0, 0, 0},
			rad:      0,
			k:        4,
			expected: 1,
		},
		{
			name:     "Large radius",
			stations: []int{1, 0, 0, 0, 1},
			rad:      10,
			k:        10,
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.stations, tt.rad, tt.k); got != tt.expected {
				t.Errorf("maxPower() = %v, want %v", got, tt.expected)
			}
		})
	}
}
