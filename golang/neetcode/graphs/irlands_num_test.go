package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	// Test nil grid
	t.Run("nil grid", func(t *testing.T) {
		result := numIslands(nil)
		assert.Equal(t, 0, result)
	})

	// Test empty grid
	t.Run("empty grid", func(t *testing.T) {
		result := numIslands([][]byte{})
		assert.Equal(t, 0, result)
	})

	// Test single cell grids
	t.Run("single cell with water", func(t *testing.T) {
		result := numIslands([][]byte{{'0'}})
		assert.Equal(t, 0, result)
	})

	t.Run("single cell with land", func(t *testing.T) {
		result := numIslands([][]byte{{'1'}})
		assert.Equal(t, 1, result)
	})

	// Test uniform grids
	t.Run("all water grid", func(t *testing.T) {
		result := numIslands([][]byte{{'0', '0'}, {'0', '0'}})
		assert.Equal(t, 0, result)
	})

	t.Run("all land grid", func(t *testing.T) {
		result := numIslands([][]byte{{'1', '1'}, {'1', '1'}})
		assert.Equal(t, 1, result)
	})

	// Test single island
	t.Run("single island", func(t *testing.T) {
		grid := [][]byte{{'1', '1', '0'}, {'1', '1', '0'}, {'0', '0', '0'}}
		result := numIslands(grid)
		assert.Equal(t, 1, result)
	})

	// Test multiple separate islands
	t.Run("multiple separate islands", func(t *testing.T) {
		grid := [][]byte{{'1', '0', '1'}, {'0', '0', '0'}, {'1', '0', '1'}}
		result := numIslands(grid)
		assert.Equal(t, 4, result)
	})

	// Test complex islands pattern
	t.Run("complex islands pattern", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}
		result := numIslands(grid)
		assert.Equal(t, 3, result)
	})

	// Test diagonal islands
	t.Run("diagonal islands", func(t *testing.T) {
		grid := [][]byte{{'1', '0', '1'}, {'0', '1', '0'}, {'1', '0', '1'}}
		result := numIslands(grid)
		assert.Equal(t, 5, result)
	})

	// Test large single island
	t.Run("large single island", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
		}
		result := numIslands(grid)
		assert.Equal(t, 1, result)
	})

	// Test checkerboard pattern
	t.Run("checkerboard pattern", func(t *testing.T) {
		grid := [][]byte{
			{'1', '0', '1', '0'},
			{'0', '1', '0', '1'},
			{'1', '0', '1', '0'},
		}
		result := numIslands(grid)
		assert.Equal(t, 6, result)
	})

	// Test single row grid
	t.Run("single row grid", func(t *testing.T) {
		grid := [][]byte{{'1', '0', '1', '0', '1'}}
		result := numIslands(grid)
		assert.Equal(t, 3, result)
	})

	// Test single column grid
	t.Run("single column grid", func(t *testing.T) {
		grid := [][]byte{{'1'}, {'0'}, {'1'}, {'0'}, {'1'}}
		result := numIslands(grid)
		assert.Equal(t, 3, result)
	})

	// Test grid with different characters
	t.Run("grid with different characters", func(t *testing.T) {
		grid := [][]byte{{'1', '2'}, {'3', '1'}} // Only '1' should be considered as land
		result := numIslands(grid)
		assert.Equal(t, 2, result) // Two separate '1's
	})
}