package trie

func minExtraChar(s string, dictionary []string) int {
	if len(s) == 0 {
		return 0
	}
	root := NewTrieNode()
	add := func(w string) {
		p := root
		for _, ch := range []byte(w) {
			if _, ok := p.nxt[ch]; !ok {
				p.nxt[ch] = NewTrieNode()
			}
			p = p.nxt[ch]
		}
		p.end = true
	}
	for _, w := range dictionary {
		add(w)
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1] + 1
		for p, j := root, i; p != nil && j < n; j++ {
			p = p.nxt[s[j]]
			if p != nil && p.end {
				dp[i] = min(dp[i], dp[j+1])
			}
		}
	}
	return dp[0]
}
