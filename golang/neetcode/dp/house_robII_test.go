package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobII(t *testing.T) {
	t.Run("empty houses", func(t *testing.T) {
		result := robII([]int{})
		assert.Equal(t, 0, result) // No houses to rob
	})

	t.Run("single house", func(t *testing.T) {
		result := robII([]int{5})
		assert.Equal(t, 5, result) // Only one house, rob it
	})

	t.Run("two houses", func(t *testing.T) {
		result := robII([]int{2, 3})
		assert.Equal(t, 3, result) // Rob the second house (can't rob both in circle)
	})

	t.Run("three houses circle", func(t *testing.T) {
		result := robII([]int{2, 3, 2})
		assert.Equal(t, 3, result) // Rob middle house (3), can't rob first and last
	})

	t.Run("classic circle example", func(t *testing.T) {
		result := robII([]int{1, 2, 3, 1})
		assert.Equal(t, 4, result) // Rob houses 2 and 4 (2+1=3) or 1 and 3 (1+3=4)
	})

	t.Run("all same amount circle", func(t *testing.T) {
		result := robII([]int{5, 5, 5, 5, 5})
		assert.Equal(t, 10, result) // Rob every other house (5+5+5=15), skip first/last conflict
	})

	t.Run("alternating pattern circle", func(t *testing.T) {
		result := robII([]int{1, 100, 1, 100, 1})
		assert.Equal(t, 200, result) // Rob houses 2 and 4 (100+100=200)
	})

	t.Run("large values circle", func(t *testing.T) {
		result := robII([]int{100, 1, 1, 100})
		assert.Equal(t, 101, result) // Rob houses 1 and 3 (100+1=101) or 2 and 4 (1+100=101)
	})

	t.Run("edge case with zeros circle", func(t *testing.T) {
		result := robII([]int{0, 0, 0, 0, 0})
		assert.Equal(t, 0, result) // All houses have zero money
	})

	t.Run("four houses circle", func(t *testing.T) {
		result := robII([]int{1, 3, 1, 3})
		assert.Equal(t, 6, result) // Rob houses 2 and 4 (3+3=6)
	})

	t.Run("five houses circle", func(t *testing.T) {
		result := robII([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 8, result) // Rob houses 2 and 4 (2+4=6) or 3 and 5 (3+5=8)
	})
}
