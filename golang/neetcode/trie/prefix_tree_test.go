package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixTree(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			op   string
			word string
			want bool
		}
	}{
		{
			name: "basic insertion and search",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "apple", false},
				{"search", "apple", true},
				{"search", "app", false},
				{"startsWith", "app", true},
				{"insert", "app", false},
				{"search", "app", true},
			},
		},
		{
			name: "empty string operations",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "", false},
				{"search", "", true},
				{"startsWith", "", true},
			},
		},
		{
			name: "multiple words with common prefixes",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "cat", false},
				{"insert", "car", false},
				{"insert", "card", false},
				{"search", "cat", true},
				{"search", "car", true},
				{"search", "card", true},
				{"search", "ca", false},
				{"startsWith", "ca", true},
				{"startsWith", "cat", true},
				{"startsWith", "car", true},
				{"startsWith", "card", true},
				{"startsWith", "cart", false},
			},
		},
		{
			name: "non-existent words",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "hello", false},
				{"search", "hell", false},
				{"search", "hello!", false},
				{"search", "world", false},
				{"startsWith", "he", true},
				{"startsWith", "hel", true},
				{"startsWith", "hello!", false},
				{"startsWith", "w", false},
			},
		},
		{
			name: "case sensitivity",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "Hello", false},
				{"search", "hello", false},
				{"search", "Hello", true},
				{"startsWith", "he", false},
				{"startsWith", "He", true},
			},
		},
		{
			name: "single character words",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "a", false},
				{"insert", "b", false},
				{"search", "a", true},
				{"search", "b", true},
				{"search", "c", false},
				{"startsWith", "a", true},
				{"startsWith", "b", true},
				{"startsWith", "c", false},
			},
		},
		{
			name: "overlapping words",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "abc", false},
				{"insert", "abcd", false},
				{"insert", "abcde", false},
				{"search", "abc", true},
				{"search", "abcd", true},
				{"search", "abcde", true},
				{"search", "ab", false},
				{"startsWith", "ab", true},
				{"startsWith", "abc", true},
				{"startsWith", "abcd", true},
				{"startsWith", "abcde", true},
				{"startsWith", "abcdef", false},
			},
		},
		{
			name: "duplicate insertions",
			operations: []struct {
				op   string
				word string
				want bool
			}{
				{"insert", "test", false},
				{"insert", "test", false}, // Duplicate insertion
				{"search", "test", true},
				{"startsWith", "test", true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()

			for _, op := range tt.operations {
				switch op.op {
				case "insert":
					trie.Insert(op.word)
				case "search":
					result := trie.Search(op.word)
					assert.Equal(t, op.want, result,
						"Search(%s) should be %v", op.word, op.want)
				case "startsWith":
					result := trie.StartsWith(op.word)
					assert.Equal(t, op.want, result,
						"StartsWith(%s) should be %v", op.word, op.want)
				}
			}
		})
	}
}
