package graphs

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	ind := make(map[string]int)
	for i, w := range wordList {
		ind[w] = i + 1
	}
	st, ed := ind[beginWord], ind[endWord]
	if ed == 0 {
		return 0
	}
	if st == 0 {
		wordList = append(wordList, beginWord)
		ind[beginWord] = len(wordList)
		st = len(wordList)
	}
	n := len(wordList)
	nxt := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			cnt := 0
			a, b := wordList[i-1], wordList[j-1]
			for k, _ := range a {
				if a[k] != b[k] {
					cnt++
					if cnt > 1 {
						break
					}
				}
			}
			if cnt == 1 {
				nxt[i] = append(nxt[i], j)
				nxt[j] = append(nxt[j], i)
			}
		}
	}

	vis := make([]bool, n+1)
	vis[st] = true
	que := [][]int{{st, 1}}
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		i, d := p[0], p[1]+1
		for _, j := range nxt[i] {
			if vis[j] {
				continue
			}
			if j == ed {
				return d
			}
			vis[j] = true
			que = append(que, []int{j, d})
		}
	}
	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int { //  O(n*m^2) instead of O(n^2*m)
	if beginWord == endWord {
		return 1
	}

	genPatterns := func(w string) (ps []string) {
		for i, _ := range w {
			ps = append(ps, w[:i]+"*"+w[i+1:])
		}
		return ps
	}

	pattern := make(map[string][]string)
	for _, w := range wordList {
		for _, p := range genPatterns(w) {
			pattern[p] = append(pattern[p], w)
		}
	}

	vis := make(map[string]bool)
	vis[beginWord] = true
	que, dis := []string{beginWord}, 1
	for len(que) > 0 {
		dis++
		sz := len(que)
		for i := 0; i < sz; i++ {
			w := que[0]
			que = que[1:]
			for _, pat := range genPatterns(w) {
				if vis[pat] {
					continue
				}
				vis[pat] = true
				for _, nxt := range pattern[pat] {
					if vis[nxt] {
						continue
					}
					if nxt == endWord {
						return dis
					}
					vis[nxt] = true
					que = append(que, nxt)
				}
			}
		}
	}
	return 0
}
