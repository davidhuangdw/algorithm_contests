package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeignDictionary(t *testing.T) {
	t.Run("empty word list", func(t *testing.T) {
		result := foreignDictionary([]string{})
		assert.Equal(t, "", result)
	})

	t.Run("single word", func(t *testing.T) {
		result := foreignDictionary([]string{"abc"})
		// Should return all unique characters in order
		assert.Contains(t, []string{"abc", "cba", "acb", "bca", "bac", "cab"}, result)
	})

	t.Run("two words simple order", func(t *testing.T) {
		words := []string{"wrt", "wrf"}
		result := foreignDictionary(words)
		// t should come before f
		assert.True(t, len(result) > 0)
		assert.Contains(t, result, "t")
		assert.Contains(t, result, "f")
		assert.True(t, indexOf(result, "t") < indexOf(result, "f"))
	})

	t.Run("classic alien dictionary example", func(t *testing.T) {
		words := []string{"wrt", "wrf", "er", "ett", "rftt"}
		result := foreignDictionary(words)
		// Expected order: w -> e -> r -> t -> f
		assert.Equal(t, "wertf", result)
	})

	t.Run("cycle detection", func(t *testing.T) {
		words := []string{"abc", "bcd", "cda", "aab"} // Creates cycle a->b->c->a
		result := foreignDictionary(words)
		assert.Equal(t, "", result) // Should return empty string for cycle
	})

	t.Run("invalid order with prefix", func(t *testing.T) {
		words := []string{"apple", "app"} // "app" is prefix of "apple" but comes after
		result := foreignDictionary(words)
		assert.Equal(t, "", result) // Invalid ordering
	})

	t.Run("multiple valid orders", func(t *testing.T) {
		words := []string{"z", "x"}
		result := foreignDictionary(words)
		assert.Equal(t, "zx", result) // z comes before x
	})

	t.Run("three words with clear ordering", func(t *testing.T) {
		words := []string{"a", "ba", "bc"}
		result := foreignDictionary(words)
		// a comes before b, but no relation between a and c
		assert.True(t, len(result) > 0)
		assert.Contains(t, result, "a")
		assert.Contains(t, result, "b")
		assert.Contains(t, result, "c")
		assert.True(t, indexOf(result, "a") < indexOf(result, "b"))
	})

	t.Run("all same characters", func(t *testing.T) {
		words := []string{"a", "aa", "aaa"}
		result := foreignDictionary(words)
		assert.Equal(t, "a", result)
	})

	t.Run("disconnected characters", func(t *testing.T) {
		words := []string{"ab", "cd"}
		result := foreignDictionary(words)
		// No relation between groups, any order is valid
		assert.True(t, len(result) == 4)
		assert.Contains(t, result, "a")
		assert.Contains(t, result, "b")
		assert.Contains(t, result, "c")
		assert.Contains(t, result, "d")
	})

	t.Run("complex valid ordering", func(t *testing.T) {
		words := []string{"ri", "xz", "qxf", "jhd", "jdd", "i", "z", "hd"}
		result := foreignDictionary(words)
		assert.True(t, len(result) > 0)
		// Verify some key relationships
		assert.True(t, indexOf(result, "r") < indexOf(result, "i"))
		assert.True(t, indexOf(result, "r") < indexOf(result, "x"))
		assert.True(t, indexOf(result, "j") < indexOf(result, "h"))
		assert.True(t, indexOf(result, "h") < indexOf(result, "d"))
	})

	t.Run("single character words", func(t *testing.T) {
		words := []string{"a", "b", "c", "d"}
		result := foreignDictionary(words)
		// All characters should be present, order doesn't matter
		assert.True(t, len(result) == 4)
		assert.Contains(t, result, "a")
		assert.Contains(t, result, "b")
		assert.Contains(t, result, "c")
		assert.Contains(t, result, "d")
	})

}

// Helper function to find index of character in string
func indexOf(s string, char string) int {
	for i, c := range s {
		if string(c) == char {
			return i
		}
	}
	return -1
}
