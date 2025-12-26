package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []float64
	}{
		// Basic cases
		{"single number", []int{5}, []float64{5}},
		{"two numbers", []int{1, 2}, []float64{1, 1.5}},
		{"three numbers", []int{1, 3, 2}, []float64{1, 2, 2}},

		// Classic cases
		{"LeetCode example", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"sorted ascending", []int{1, 2, 3, 4, 5}, []float64{1, 1.5, 2, 2.5, 3}},
		{"sorted descending", []int{5, 4, 3, 2, 1}, []float64{5, 4.5, 4, 3.5, 3}},

		// Edge cases
		{"all same numbers", []int{2, 2, 2, 2}, []float64{2, 2, 2, 2}},
		{"negative numbers", []int{-1, -2, -3}, []float64{-1, -1.5, -2}},
		{"mixed positive negative", []int{-1, 2, -3, 4}, []float64{-1, 0.5, -1, 0.5}},
		{"large numbers", []int{1000, 2000, 3000}, []float64{1000, 1500, 2000}},

		// Complex cases
		{"alternating pattern", []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, []float64{1, 5.5, 2, 5.5, 3, 5.5, 4, 5.5, 5, 5.5}},
		{"random sequence", []int{3, 1, 4, 1, 5, 9, 2, 6}, []float64{3, 2, 3, 2, 3, 3.5, 3, 3.5}},
		{"empty case", []int{}, []float64{0.0}}, // Empty case test
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			medians := make([]float64, 0, len(tt.nums))

			for i, num := range tt.nums {
				mf.AddNum(num)
				median := mf.FindMedian()
				medians = append(medians, median)

				assert.InDelta(t, tt.expected[i], median, 1e-9,
					"Median after adding %v should be %v, got %v",
					tt.nums[:i+1], tt.expected[i], median)
			}

			// Handle empty case separately
			if len(tt.nums) == 0 {
				median := mf.FindMedian()
				assert.Equal(t, 0.0, median, "Empty median finder should return 0.0")
			}
		})
	}
}