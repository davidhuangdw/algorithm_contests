package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDecodings(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		result := numDecodings("")
		assert.Equal(t, 1, result) // Empty string has 1 way: decode as empty
	})

	t.Run("single digit valid", func(t *testing.T) {
		result := numDecodings("1")
		assert.Equal(t, 1, result) // "1" -> "A"
	})

	t.Run("single digit invalid", func(t *testing.T) {
		result := numDecodings("0")
		assert.Equal(t, 0, result) // "0" is invalid
	})

	t.Run("two digits both valid", func(t *testing.T) {
		result := numDecodings("12")
		assert.Equal(t, 2, result) // "12" -> "AB" or "L"
	})

	t.Run("two digits with zero", func(t *testing.T) {
		result := numDecodings("10")
		assert.Equal(t, 1, result) // "10" -> "J" only
	})

	t.Run("classic example", func(t *testing.T) {
		result := numDecodings("226")
		assert.Equal(t, 3, result) // "226" -> "BBF", "BZ", "VF"
	})

	t.Run("leading zero", func(t *testing.T) {
		result := numDecodings("06")
		assert.Equal(t, 0, result) // Leading zero is invalid
	})

	t.Run("all ones", func(t *testing.T) {
		result := numDecodings("111")
		assert.Equal(t, 3, result) // "111" -> "AAA", "KA", "AK"
	})

	t.Run("large number", func(t *testing.T) {
		result := numDecodings("27")
		assert.Equal(t, 1, result) // "27" -> "BG" only (27 > 26)
	})

	t.Run("complex case", func(t *testing.T) {
		result := numDecodings("1111")
		assert.Equal(t, 5, result) // "1111" -> "AAAA", "KAA", "AKA", "AAK", "KK"
	})

	t.Run("edge case with zeros", func(t *testing.T) {
		result := numDecodings("100")
		assert.Equal(t, 0, result) // "100" has invalid zeros
	})

	t.Run("valid zeros", func(t *testing.T) {
		result := numDecodings("110")
		assert.Equal(t, 1, result) // "110" -> "AJ" only
	})

	t.Run("maximum valid two-digit", func(t *testing.T) {
		result := numDecodings("26")
		assert.Equal(t, 2, result) // "26" -> "BF" or "Z"
	})

	t.Run("just above maximum", func(t *testing.T) {
		result := numDecodings("27")
		assert.Equal(t, 1, result) // "27" -> "BG" only
	})
}
