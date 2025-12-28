package trie

type WordDictionary struct {
	root *TrieNode
}

func ConstructorWd() WordDictionary {
	return WordDictionary{NewTrieNode()}
}

func (wo *WordDictionary) AddWord(word string) {
	p := wo.root
	for _, ch := range []byte(word) {
		if _, ok := p.nxt[ch]; !ok {
			p.nxt[ch] = NewTrieNode()
		}
		p = p.nxt[ch]
	}
	p.end = true
}

func (wo *WordDictionary) Search(word string) bool {
	wd := []byte(word)
	var dfs func(i int, cur *TrieNode) bool
	dfs = func(i int, cur *TrieNode) bool {
		if cur == nil {
			return false
		}
		if i >= len(wd) {
			return cur.end
		}
		if wd[i] != '.' {
			return dfs(i+1, cur.nxt[wd[i]])
		}
		for _, p := range cur.nxt {
			if dfs(i+1, p) {
				return true
			}
		}
		return false
	}
	return dfs(0, wo.root)
}
