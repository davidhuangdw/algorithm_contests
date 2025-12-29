package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	t.Run("empty houses", func(t *testing.T) {
		result := rob([]int{})
		assert.Equal(t, 0, result) // No houses to rob
	})

	t.Run("single house", func(t *testing.T) {
		result := rob([]int{5})
		assert.Equal(t, 5, result) // Only one house, rob it
	})

	t.Run("two houses", func(t *testing.T) {
		result := rob([]int{2, 3})
		assert.Equal(t, 3, result) // Rob the second house (3 > 2)
	})

	t.Run("three houses", func(t *testing.T) {
		result := rob([]int{1, 2, 3})
		assert.Equal(t, 4, result) // Rob houses 1 and 3 (1+3=4)
	})

	t.Run("classic example", func(t *testing.T) {
		result := rob([]int{2, 7, 9, 3, 1})
		assert.Equal(t, 12, result) // Rob houses 2 and 4 (7+3=10) or 1,3,5 (2+9+1=12)
	})

	t.Run("all same amount", func(t *testing.T) {
		result := rob([]int{5, 5, 5, 5, 5})
		assert.Equal(t, 15, result) // Rob every other house (5+5+5=15)
	})

	t.Run("alternating pattern", func(t *testing.T) {
		result := rob([]int{1, 100, 1, 100, 1})
		assert.Equal(t, 200, result) // Rob houses 2 and 4 (100+100=200)
	})

	t.Run("large values", func(t *testing.T) {
		result := rob([]int{100, 1, 1, 100})
		assert.Equal(t, 200, result) // Rob houses 1 and 4 (100+100=200)
	})

	t.Run("increasing values", func(t *testing.T) {
		result := rob([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 9, result) // Rob houses 2 and 4 (2+4=6) or 1,3,5 (1+3+5=9)
	})

	t.Run("edge case with zeros", func(t *testing.T) {
		result := rob([]int{0, 0, 0, 0, 0})
		assert.Equal(t, 0, result) // All houses have zero money
	})
}
