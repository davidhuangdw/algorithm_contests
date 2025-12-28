package trie

func findWords(board [][]byte, words []string) []string {
	ans := make([]string, 0)
	if len(board) == 0 || len(board[0]) == 0 {
		return ans
	}

	root := NewTrieNode()
	addWord := func(word string) {
		wd, p := []byte(word), root
		for _, ch := range wd {
			if _, ok := p.nxt[ch]; !ok {
				p.nxt[ch] = NewTrieNode()
			}
			p = p.nxt[ch]
		}
		p.end = true
	}

	for _, w := range words {
		addWord(w)
	}

	n, m, fnd := len(board), len(board[0]), make(map[string]bool)
	var dfs func(i, j int, cur *TrieNode, pre []byte)
	dfs = func(i, j int, cur *TrieNode, pre []byte) {
		if !(0 <= i && i < n && 0 <= j && j < m && board[i][j] != '#') {
			return
		}
		ch := board[i][j]
		cur, pre = cur.nxt[ch], append(pre, ch)
		if cur == nil {
			return
		}
		if cur.end {
			fnd[string(pre)] = true
		}

		board[i][j] = '#'
		dfs(i+1, j, cur, pre)
		dfs(i-1, j, cur, pre)
		dfs(i, j+1, cur, pre)
		dfs(i, j-1, cur, pre)
		board[i][j] = ch
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j, root, nil)
		}
	}
	for w, _ := range fnd {
		ans = append(ans, w)
	}
	return ans
}
