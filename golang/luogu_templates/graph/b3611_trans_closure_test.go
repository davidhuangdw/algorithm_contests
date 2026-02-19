package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransClosure(t *testing.T) {
	t.Run("Single node no edges", func(t *testing.T) {
		input := [][]int{{0}}
		expected := [][]int{{0}}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Single edge", func(t *testing.T) {
		input := [][]int{
			{0, 1},
			{0, 0},
		}
		expected := [][]int{
			{0, 1},
			{0, 0},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Chain creates transitive edge", func(t *testing.T) {
		input := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 1, 1}, // 0->1->2 creates 0->2
			{0, 0, 1},
			{0, 0, 0},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Triangle all connected", func(t *testing.T) {
		input := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 0, 0},
		}
		expected := [][]int{
			{1, 1, 1}, // cycle creates all paths
			{1, 1, 1},
			{1, 1, 1},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Disconnected components", func(t *testing.T) {
		input := [][]int{
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 0, 0},
		}
		expected := [][]int{
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 0, 0},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Already transitive", func(t *testing.T) {
		input := [][]int{
			{0, 1, 1},
			{0, 0, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 1, 1},
			{0, 0, 1},
			{0, 0, 0},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Self-loop propagates", func(t *testing.T) {
		input := [][]int{
			{0, 1, 0},
			{0, 1, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 1, 1}, // 0->1 and 1 has self-loop and 1->2
			{0, 1, 1}, // 1 has self-loop and 1->2
			{0, 0, 0},
		}
		result := transClosure(input)
		assert.Equal(t, expected, result)
	})
}
