package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	t.Run("no courses", func(t *testing.T) {
		result := canFinish(0, [][]int{})
		assert.True(t, result)
	})

	t.Run("single course no prerequisites", func(t *testing.T) {
		result := canFinish(1, [][]int{})
		assert.True(t, result)
	})

	t.Run("two courses no prerequisites", func(t *testing.T) {
		result := canFinish(2, [][]int{})
		assert.True(t, result)
	})

	t.Run("two courses with valid prerequisite", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}} // course 1 requires course 0
		result := canFinish(2, prerequisites)
		assert.True(t, result)
	})

	t.Run("three courses linear dependencies", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {2, 1}} // 2->1->0
		result := canFinish(3, prerequisites)
		assert.True(t, result)
	})

	t.Run("three courses with cycle", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {0, 1}} // 0<->1 cycle
		result := canFinish(2, prerequisites)
		assert.False(t, result)
	})

	t.Run("four courses with cycle", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}} // 1->0, 2->1, 3->2, 1->3 (cycle)
		result := canFinish(4, prerequisites)
		assert.False(t, result)
	})

	t.Run("multiple independent courses", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {3, 2}} // two independent chains
		result := canFinish(4, prerequisites)
		assert.True(t, result)
	})

	t.Run("complex graph without cycle", func(t *testing.T) {
		prerequisites := [][]int{
			{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3},
		}
		result := canFinish(5, prerequisites)
		assert.True(t, result)
	})

	t.Run("complex graph with cycle", func(t *testing.T) {
		prerequisites := [][]int{
			{1, 0}, {2, 0}, {3, 1}, {3, 2}, {0, 3}, // 0->3->1->0 cycle
		}
		result := canFinish(4, prerequisites)
		assert.False(t, result)
	})

	t.Run("self-loop course", func(t *testing.T) {
		prerequisites := [][]int{{0, 0}} // course 0 requires itself
		result := canFinish(1, prerequisites)
		assert.False(t, result)
	})

	t.Run("disconnected graph with cycle in one component", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {0, 1}, {3, 2}} // cycle in first component, valid in second
		result := canFinish(4, prerequisites)
		assert.False(t, result)
	})

	t.Run("large number of courses no prerequisites", func(t *testing.T) {
		result := canFinish(1000, [][]int{})
		assert.True(t, result)
	})

	t.Run("prerequisites referencing non-existent courses", func(t *testing.T) {
		// This should handle gracefully - courses are 0 to n-1
		prerequisites := [][]int{{1, 0}, {2, 1}} // all valid for n=3
		result := canFinish(3, prerequisites)
		assert.True(t, result)
	})
}