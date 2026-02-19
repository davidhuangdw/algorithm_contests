package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isPrefixFree checks that no code is a prefix of another
func isPrefixFree(codes []string) bool {
	for i, a := range codes {
		for j, b := range codes {
			if i != j && strings.HasPrefix(b, a) {
				return false
			}
		}
	}
	return true
}

func wpl(freq []int, codes []string) int {
	total := 0
	for i, c := range codes {
		total += freq[i] * len(c)
	}
	return total
}

func TestHuffmanCode(t *testing.T) {
	t.Run("Single symbol", func(t *testing.T) {
		codes := HuffmanCode([]int{5})
		assert.Equal(t, []string{"0"}, codes)
	})

	t.Run("Two symbols", func(t *testing.T) {
		codes := HuffmanCode([]int{3, 5})
		assert.ElementsMatch(t, []string{"0", "1"}, codes)
	})

	t.Run("Prefix-free", func(t *testing.T) {
		codes := HuffmanCode([]int{1, 1, 2, 3, 5})
		assert.True(t, isPrefixFree(codes))
	})

	t.Run("Higher freq gets shorter or equal code", func(t *testing.T) {
		freq := []int{1, 2, 4, 8}
		codes := HuffmanCode(freq)
		// each freq is strictly larger, so lengths must be non-increasing
		for i := 0; i < len(freq)-1; i++ {
			assert.LessOrEqual(t, len(codes[i+1]), len(codes[i]),
				"freq[%d]=%d should have code len <= freq[%d]=%d", i+1, freq[i+1], i, freq[i])
		}
	})

	t.Run("Optimal WPL", func(t *testing.T) {
		freq := []int{1, 1, 2, 3}
		codes := HuffmanCode(freq)
		assert.Equal(t, 13, wpl(freq, codes))
	})
}
