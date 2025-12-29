package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		result := longestPalindrome("")
		assert.Equal(t, "", result) // Empty string should return empty
	})

	t.Run("single character", func(t *testing.T) {
		result := longestPalindrome("a")
		assert.Equal(t, "a", result) // Single character is always palindrome
	})

	t.Run("two characters palindrome", func(t *testing.T) {
		result := longestPalindrome("aa")
		assert.Equal(t, "aa", result) // Two same characters form palindrome
	})

	t.Run("two characters not palindrome", func(t *testing.T) {
		result := longestPalindrome("ab")
		assert.Equal(t, "a", result) // Should return first character
	})

	t.Run("classic example", func(t *testing.T) {
		result := longestPalindrome("babad")
		assert.Contains(t, []string{"bab", "aba"}, result) // Both are valid longest palindromes
	})

	t.Run("even length palindrome", func(t *testing.T) {
		result := longestPalindrome("cbbd")
		assert.Equal(t, "bb", result) // Longest palindrome is "bb"
	})

	t.Run("full palindrome", func(t *testing.T) {
		result := longestPalindrome("racecar")
		assert.Equal(t, "racecar", result) // Entire string is palindrome
	})

	t.Run("multiple palindromes", func(t *testing.T) {
		result := longestPalindrome("abacdfgdcaba")
		assert.Equal(t, "aba", result) // Longest palindrome is "aba"
	})

	t.Run("all same characters", func(t *testing.T) {
		result := longestPalindrome("aaaa")
		assert.Equal(t, "aaaa", result) // Entire string is palindrome
	})

	t.Run("numeric palindrome", func(t *testing.T) {
		result := longestPalindrome("12321")
		assert.Equal(t, "12321", result) // Numeric palindrome
	})

}
