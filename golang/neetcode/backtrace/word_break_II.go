package backtrace

import "strings"

func wordBreak1(s string, wordDict []string) []string {
	rs := []rune(s)
	root, ans, n := NewTrieNode(), make([]string, 0), len(rs)
	for _, word := range wordDict {
		root.AddWord(word)
	}

	var dfs func(i int, pre []string, cur []rune, nd *TrieNode)
	dfs = func(i int, pre []string, cur []rune, nd *TrieNode) {
		if nd.end {
			dfs(i, append(pre, string(cur)), nil, root)
		}
		if i == n {
			if len(cur) == 0 {
				ans = append(ans, strings.Join(pre, " "))
			}
			return
		}
		if p, ok := nd.nxt[rs[i]]; ok {
			dfs(i+1, pre, append(cur, rs[i]), p)
		}
	}
	dfs(0, nil, nil, root)
	return ans
}

func wordBreak(s string, wordDict []string) []string { // by dp
	rs := []rune(s)
	root, n := NewTrieNode(), len(rs)
	for _, word := range wordDict {
		root.AddWord(word)
	}
	dp := make(map[int][]string)
	dp[n] = []string{""}
	var dfs func(i int) []string
	dfs = func(i int) []string {
		if _, ok := dp[i]; !ok {
			res := make([]string, 0)
			p := root
			for j := i; j < n; j++ {
				p = p.nxt[rs[j]]
				if p == nil {
					break
				}
				if !p.end {
					continue
				}
				first := string(rs[i : j+1])
				for _, remain := range dfs(j + 1) {
					if remain == "" {
						res = append(res, first)
					} else {
						res = append(res, first+" "+remain)
					}
				}
			}
			dp[i] = res
		}
		return dp[i]
	}
	return dfs(0)
}

type TrieNode struct {
	nxt map[rune]*TrieNode
	end bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{nxt: make(map[rune]*TrieNode)}
}

func (tr *TrieNode) AddWord(wd string) {
	p := tr
	for _, ch := range wd {
		if _, ok := p.nxt[ch]; !ok {
			p.nxt[ch] = NewTrieNode()
		}
		p = p.nxt[ch]
	}
	p.end = true
}
