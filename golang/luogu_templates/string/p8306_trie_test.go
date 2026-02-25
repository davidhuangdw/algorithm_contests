package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	// Test basic functionality
	trie := NewTrie()

	// Add words
	trie.Add("apple")
	trie.Add("app")
	trie.Add("application")
	trie.Add("banana")
	trie.Add("app") // Add duplicate word

	// Test prefix counts
	assert.Equal(t, 4, trie.Count("app"), "Prefix 'app' should match 4 words (including duplicate)")
	assert.Equal(t, 1, trie.Count("banana"), "Prefix 'banana' should match 1 word")
	assert.Equal(t, 0, trie.Count("orange"), "Prefix 'orange' should match 0 words")
	assert.Equal(t, 2, trie.Count("appl"), "Prefix 'appl' should match 2 words")

	// Test empty trie
	emptyTrie := NewTrie()
	assert.Equal(t, 0, emptyTrie.Count("any"), "Empty trie should return 0 for any prefix")

	// Test with single character words
	singleCharTrie := NewTrie()
	singleCharTrie.Add("a")
	singleCharTrie.Add("b")
	singleCharTrie.Add("a") // Add duplicate
	assert.Equal(t, 2, singleCharTrie.Count("a"), "Prefix 'a' should match 2 words")
	assert.Equal(t, 1, singleCharTrie.Count("b"), "Prefix 'b' should match 1 word")

	// Test with empty string
	emptyStrTrie := NewTrie()
	emptyStrTrie.Add("")
	emptyStrTrie.Add("")
	assert.Equal(t, 2, emptyStrTrie.Count(""), "Empty string prefix should match 2 words")

	// Test with words that are prefixes of others
	prefixTrie := NewTrie()
	prefixTrie.Add("a")
	prefixTrie.Add("ab")
	prefixTrie.Add("abc")
	assert.Equal(t, 3, prefixTrie.Count("a"), "Prefix 'a' should match 3 words")
	assert.Equal(t, 2, prefixTrie.Count("ab"), "Prefix 'ab' should match 2 words")
	assert.Equal(t, 1, prefixTrie.Count("abc"), "Prefix 'abc' should match 1 word")
}
