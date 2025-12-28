package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacificAtlantic(t *testing.T) {
	t.Run("empty grid", func(t *testing.T) {
		result := pacificAtlantic([][]int{})
		assert.Empty(t, result)
	})

	t.Run("single cell grid", func(t *testing.T) {
		heights := [][]int{{5}}
		result := pacificAtlantic(heights)
		assert.Equal(t, [][]int{{0, 0}}, result)
	})

	t.Run("1x2 grid", func(t *testing.T) {
		heights := [][]int{{1, 2}}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 0}, {0, 1}}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("2x1 grid", func(t *testing.T) {
		heights := [][]int{{1}, {2}}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 0}, {1, 0}}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("2x2 grid with valley", func(t *testing.T) {
		heights := [][]int{
			{1, 2},
			{2, 1},
		}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 1}, {1, 0}}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("3x3 grid with mountain", func(t *testing.T) {
		heights := [][]int{
			{1, 2, 3},
			{2, 3, 2},
			{3, 2, 1},
		}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 2}, {1, 1}, {2, 0}}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("complex 5x5 grid", func(t *testing.T) {
		heights := [][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}
		result := pacificAtlantic(heights)
		expected := [][]int{
			{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
		}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("all same height", func(t *testing.T) {
		heights := [][]int{
			{3, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}
		result := pacificAtlantic(heights)
		expected := [][]int{
			{0, 0}, {0, 1}, {0, 2},
			{1, 0}, {1, 1}, {1, 2},
			{2, 0}, {2, 1}, {2, 2},
		}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("descending from center", func(t *testing.T) {
		heights := [][]int{
			{1, 1, 1},
			{1, 2, 1},
			{1, 1, 1},
		}
		result := pacificAtlantic(heights)
		expected := [][]int{
			{0, 0}, {0, 1}, {0, 2},
			{1, 0}, {1, 1}, {1, 2},
			{2, 0}, {2, 1}, {2, 2},
		}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("single row grid", func(t *testing.T) {
		heights := [][]int{{1, 2, 3, 4, 5}}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("single column grid", func(t *testing.T) {
		heights := [][]int{{1}, {2}, {3}, {4}, {5}}
		result := pacificAtlantic(heights)
		expected := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}
		assert.ElementsMatch(t, expected, result)
	})
}
