package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFindSet(t *testing.T) {
	t.Run("Initially all separate", func(t *testing.T) {
		uf := NewUnionFindSet(5)
		for i := 1; i <= 5; i++ {
			for j := i + 1; j <= 5; j++ {
				assert.NotEqual(t, uf.find(i), uf.find(j))
			}
		}
	})

	t.Run("Union merges components", func(t *testing.T) {
		uf := NewUnionFindSet(5)
		uf.union(1, 2)
		uf.union(3, 4)
		assert.Equal(t, uf.find(1), uf.find(2))
		assert.Equal(t, uf.find(3), uf.find(4))
		assert.NotEqual(t, uf.find(1), uf.find(3))
	})

	t.Run("Transitive union", func(t *testing.T) {
		uf := NewUnionFindSet(4)
		uf.union(1, 2)
		uf.union(2, 3)
		assert.Equal(t, uf.find(1), uf.find(3))
		assert.NotEqual(t, uf.find(1), uf.find(4))
	})

	t.Run("Union with self is no-op", func(t *testing.T) {
		uf := NewUnionFindSet(3)
		uf.union(2, 2)
		assert.NotEqual(t, uf.find(1), uf.find(2))
	})

	t.Run("Union already same component", func(t *testing.T) {
		uf := NewUnionFindSet(3)
		uf.union(1, 2)
		uf.union(1, 2) // idempotent
		assert.Equal(t, uf.find(1), uf.find(2))
	})

	t.Run("All merged", func(t *testing.T) {
		uf := NewUnionFindSet(4)
		uf.union(1, 2)
		uf.union(2, 3)
		uf.union(3, 4)
		root := uf.find(1)
		for i := 2; i <= 4; i++ {
			assert.Equal(t, root, uf.find(i))
		}
	})
}
