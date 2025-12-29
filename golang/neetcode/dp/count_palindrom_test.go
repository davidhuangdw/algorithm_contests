package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstrings(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		result := countSubstrings("")
		assert.Equal(t, 0, result) // No palindromic substrings in empty string
	})

	t.Run("single character", func(t *testing.T) {
		result := countSubstrings("a")
		assert.Equal(t, 1, result) // Single character: "a"
	})

	t.Run("two same characters", func(t *testing.T) {
		result := countSubstrings("aa")
		assert.Equal(t, 3, result) // "a", "a", "aa"
	})

	t.Run("two different characters", func(t *testing.T) {
		result := countSubstrings("ab")
		assert.Equal(t, 2, result) // "a", "b"
	})

	t.Run("three characters palindrome", func(t *testing.T) {
		result := countSubstrings("aaa")
		assert.Equal(t, 6, result) // "a", "a", "a", "aa", "aa", "aaa"
	})

	t.Run("classic example", func(t *testing.T) {
		result := countSubstrings("abc")
		assert.Equal(t, 3, result) // "a", "b", "c"
	})

	t.Run("mixed palindrome", func(t *testing.T) {
		result := countSubstrings("aba")
		assert.Equal(t, 4, result) // "a", "b", "a", "aba"
	})

	t.Run("longer palindrome", func(t *testing.T) {
		result := countSubstrings("abba")
		assert.Equal(t, 6, result) // "a", "b", "b", "a", "bb", "abba"
	})

	t.Run("all same characters", func(t *testing.T) {
		result := countSubstrings("aaaa")
		assert.Equal(t, 10, result) // 4 singles + 3 doubles + 2 triples + 1 quadruple = 10
	})

	t.Run("numeric palindrome", func(t *testing.T) {
		result := countSubstrings("121")
		assert.Equal(t, 4, result) // "1", "2", "1", "121"
	})

	t.Run("special characters", func(t *testing.T) {
		result := countSubstrings("a!a")
		assert.Equal(t, 4, result) // "a", "!", "a", "a!a"
	})

	t.Run("complex palindrome", func(t *testing.T) {
		result := countSubstrings("racecar")
		assert.Equal(t, 10, result) // Includes "r", "a", "c", "e", "c", "a", "r", "cec", "aceca", "racecar"
	})
}
