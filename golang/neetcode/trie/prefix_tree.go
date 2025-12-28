package trie

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{NewTrieNode()}
}

func (pr *PrefixTree) Insert(word string) {
	nd := pr.root
	for _, ch := range []byte(word) {
		if _, ok := nd.nxt[ch]; !ok {
			nd.nxt[ch] = NewTrieNode()
		}
		nd = nd.nxt[ch]
	}
	nd.end = true
}

func (pr *PrefixTree) Search(word string) bool {
	nd := pr.walk(word)
	return nd != nil && nd.end
}

func (pr *PrefixTree) StartsWith(prefix string) bool {
	nd := pr.walk(prefix)
	return nd != nil

}

func (pr *PrefixTree) walk(prefix string) *TrieNode {
	nd := pr.root
	for _, ch := range []byte(prefix) {
		nd = nd.nxt[ch]
		if nd == nil {
			break
		}
	}
	return nd
}
