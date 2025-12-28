package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	// Test constructor
	t.Run("ConstructorWd creates empty dictionary", func(t *testing.T) {
		dict := ConstructorWd()
		assert.NotNil(t, dict.root)
		assert.NotNil(t, dict.root.nxt)
		assert.Empty(t, dict.root.nxt)
		assert.False(t, dict.root.end)
	})

	// Test edge cases
	t.Run("Edge cases", func(t *testing.T) {
		dict := ConstructorWd()

		// Test search on empty dictionary
		assert.False(t, dict.Search("test"))
		assert.False(t, dict.Search("."))
		assert.False(t, dict.Search(".."))

		// Add empty string and test
		dict.AddWord("")
		assert.True(t, dict.Search(""))
		assert.False(t, dict.Search("."))

		// Test with single character
		dict.AddWord("a")
		assert.True(t, dict.Search("a"))
		assert.True(t, dict.Search("."))
		assert.False(t, dict.Search("b"))
		assert.False(t, dict.Search(".."))
	})

	// Test basic operations without dots
	t.Run("Basic operations without dots", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("bad")
		dict.AddWord("dad")
		dict.AddWord("mad")

		assert.False(t, dict.Search("pad"))
		assert.True(t, dict.Search("bad"))
		assert.True(t, dict.Search(".ad"))
		assert.True(t, dict.Search("b.."))
	})

	// Test single character words
	t.Run("Single character words", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("a")
		dict.AddWord("b")

		assert.True(t, dict.Search("a"))
		assert.True(t, dict.Search("b"))
		assert.False(t, dict.Search("c"))
		assert.True(t, dict.Search("."))
		assert.False(t, dict.Search(".."))
	})

	// Test empty string operations
	t.Run("Empty string operations", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("")

		assert.True(t, dict.Search(""))
		assert.False(t, dict.Search("."))
	})

	// Test multiple dots in pattern
	t.Run("Multiple dots in pattern", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("hello")
		dict.AddWord("world")

		assert.True(t, dict.Search("h.llo"))
		assert.True(t, dict.Search("he.lo"))
		assert.True(t, dict.Search("hel.o"))
		assert.True(t, dict.Search("hell."))
		assert.True(t, dict.Search("w.rld"))
		assert.True(t, dict.Search("wo.ld"))
		assert.True(t, dict.Search("wor.d"))
		assert.True(t, dict.Search("worl."))
		assert.True(t, dict.Search("....."))
		assert.True(t, dict.Search("h...."))
		assert.True(t, dict.Search("....o"))
		assert.True(t, dict.Search("..l.."))
		assert.False(t, dict.Search("......"))
	})

	// Test partial matches and non-matches
	t.Run("Partial matches and non-matches", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("cat")
		dict.AddWord("car")
		dict.AddWord("card")

		assert.True(t, dict.Search("cat"))
		assert.True(t, dict.Search("car"))
		assert.True(t, dict.Search("card"))
		assert.False(t, dict.Search("ca"))
		assert.True(t, dict.Search("c.t"))
		assert.True(t, dict.Search("c.r"))
		assert.True(t, dict.Search("c.rd"))
		assert.True(t, dict.Search("c..d"))
		assert.True(t, dict.Search("c..."))
		assert.True(t, dict.Search("...."))
		assert.False(t, dict.Search("....."))
		assert.False(t, dict.Search("bat"))
		assert.False(t, dict.Search("b.t"))
	})

	// Test duplicate words
	t.Run("Duplicate words", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("test")
		dict.AddWord("test") // Duplicate insertion

		assert.True(t, dict.Search("test"))
		assert.True(t, dict.Search("t.st"))
		assert.True(t, dict.Search(".est"))
		assert.True(t, dict.Search("te.t"))
		assert.True(t, dict.Search("tes."))
	})

	// Test case sensitivity
	t.Run("Case sensitivity", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("Hello")

		assert.False(t, dict.Search("hello"))
		assert.True(t, dict.Search("Hello"))
		assert.True(t, dict.Search("H.llo"))
		assert.False(t, dict.Search("h.llo"))
		assert.True(t, dict.Search("Hel.o"))
	})

	// Test complex dot patterns
	t.Run("Complex dot patterns", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("abcdef")

		assert.True(t, dict.Search("a.c.e."))
		assert.True(t, dict.Search(".b.d.f"))
		assert.True(t, dict.Search("..c..f"))
		assert.True(t, dict.Search("a..d.f"))
		assert.True(t, dict.Search("a.c..f"))
		assert.True(t, dict.Search("......"))
		assert.False(t, dict.Search("......."))
		assert.False(t, dict.Search("abcde"))
		assert.False(t, dict.Search("abcdefg"))
	})

	// Test multiple words with overlapping patterns
	t.Run("Multiple words with overlapping patterns", func(t *testing.T) {
		dict := ConstructorWd()
		dict.AddWord("abc")
		dict.AddWord("abd")
		dict.AddWord("abe")

		assert.True(t, dict.Search("ab."))
		assert.True(t, dict.Search("a.c"))
		assert.True(t, dict.Search("a.d"))
		assert.True(t, dict.Search("a.e"))
		assert.True(t, dict.Search("..."))
		assert.False(t, dict.Search("abf"))
		assert.False(t, dict.Search("a.f"))
	})
}
