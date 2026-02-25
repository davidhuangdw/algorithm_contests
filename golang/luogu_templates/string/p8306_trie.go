package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type TrieNode struct {
	nxt map[byte]*TrieNode
	cnt int
}
type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{nxt: make(map[byte]*TrieNode)}
}
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (tr *Trie) Add(word string) {
	p := tr.root
	p.cnt++
	for _, ch := range []byte(word) {
		if p.nxt[ch] == nil {
			p.nxt[ch] = NewTrieNode()
		}
		p = p.nxt[ch]
		p.cnt++
	}
}

func (tr *Trie) Count(pre string) int {
	p := tr.root
	for _, ch := range []byte(pre) {
		p = p.nxt[ch]
		if p == nil {
			return 0
		}
	}
	return p.cnt
}

func P8306_trie(IN io.Reader, OUT io.Writer) {
	var T, n, q int
	fmt.Fscan(IN, &T)
	for ; T > 0; T-- {
		fmt.Fscan(IN, &n, &q)
		tr := NewTrie()

		var s string
		for i := 0; i < n; i++ {
			fmt.Fscan(IN, &s)
			tr.Add(s)
		}

		for i := 0; i < q; i++ {
			fmt.Fscan(IN, &s)
			fmt.Fprintln(OUT, tr.Count(s))
		}
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P8306_trie(IN, OUT)
}
