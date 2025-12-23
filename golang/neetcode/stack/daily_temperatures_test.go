package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		t    []int
		want []int
	}{
		// Basic cases
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
		{[]int{90, 80, 70, 60}, []int{0, 0, 0, 0}},
		{[]int{70, 70, 70, 70}, []int{0, 0, 0, 0}}, // All same temperatures - no strictly greater

		// Edge cases
		{[]int{}, []int{}},
		{[]int{100}, []int{0}},
		{[]int{0}, []int{0}},
		{[]int{-10}, []int{0}},
		{[]int{100, 90}, []int{0, 0}},
		{[]int{90, 100}, []int{1, 0}},

		// Complex patterns
		{[]int{73, 72, 71, 74}, []int{3, 2, 1, 0}},
		{[]int{30, 31, 25, 30, 35}, []int{1, 3, 1, 1, 0}},
		{[]int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}, []int{3, 1, 1, 2, 1, 1, 0, 1, 1, 0}},
		{[]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}, []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}},

		// All increasing
		{[]int{10, 20, 30, 40, 50}, []int{1, 1, 1, 1, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 1, 1, 1, 0}},

		// All decreasing
		{[]int{50, 40, 30, 20, 10}, []int{0, 0, 0, 0, 0}},
		{[]int{5, 4, 3, 2, 1}, []int{0, 0, 0, 0, 0}},

		// Mixed patterns
		{[]int{70, 69, 68, 70, 69, 68, 70}, []int{0, 2, 1, 0, 2, 1, 0}},
		{[]int{10, 15, 12, 18, 14, 20}, []int{1, 2, 1, 2, 1, 0}},
		{[]int{25, 30, 20, 35, 15, 40}, []int{1, 2, 1, 2, 1, 0}},

		// Large temperature ranges
		{[]int{-10, 0, 10, 20, 30}, []int{1, 1, 1, 1, 0}},
		{[]int{100, 95, 90, 85, 80}, []int{0, 0, 0, 0, 0}},
		{[]int{0, 100, 0, 100, 0}, []int{1, 0, 1, 0, 0}},

		// Long sequences
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{5, 10, 5, 10, 5, 10, 5, 10, 5, 10}, []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}},

		// Real-world temperature patterns
		{[]int{65, 70, 68, 72, 75, 73, 78, 76, 80}, []int{1, 2, 1, 1, 2, 1, 2, 1, 0}},
		{[]int{32, 35, 33, 40, 38, 42, 39, 45}, []int{1, 2, 1, 2, 1, 2, 1, 0}},
		{[]int{20, 22, 18, 25, 23, 28, 26, 30}, []int{1, 2, 1, 2, 1, 2, 1, 0}},

		// Extreme values
		{[]int{-273, 1000}, []int{1, 0}},
		{[]int{1000, -273}, []int{0, 0}},
		{[]int{-100, -50, 0, 50, 100}, []int{1, 1, 1, 1, 0}},

		// Single temperature patterns
		{[]int{50, 50, 50, 51, 50, 50}, []int{3, 2, 1, 0, 0, 0}},
		{[]int{30, 30, 31, 30, 30}, []int{2, 1, 0, 0, 0}},
		{[]int{25, 25, 25, 25}, []int{0, 0, 0, 0}},

		// Complex temperature fluctuations
		{[]int{70, 75, 65, 80, 60, 85, 55, 90}, []int{1, 2, 1, 2, 1, 2, 1, 0}},
		{[]int{40, 45, 35, 50, 30, 55, 25, 60}, []int{1, 2, 1, 2, 1, 2, 1, 0}},
		{[]int{10, 20, 5, 25, 0, 30, -5, 35}, []int{1, 2, 1, 2, 1, 2, 1, 0}},

		// Large arrays
		{make([]int, 100), make([]int, 100)}, // All zeros for constant temperatures
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := dailyTemperatures(tt.t)
			assert.Equal(t, tt.want, result,
				"dailyTemperatures(%v) = %v, want %v", tt.t, result, tt.want)
		})
	}
}
