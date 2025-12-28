package trie

type TrieNode struct {
	nxt map[byte]*TrieNode
	end bool
}

func NewTrieNode() *TrieNode { return &TrieNode{nxt: make(map[byte]*TrieNode)} }
