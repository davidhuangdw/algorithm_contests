package graphs

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	if len(edges) != n-1 {
		return false
	}
	nxt := make([][]int, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		if i == j {
			return false
		}
		nxt[i] = append(nxt[i], j)
		nxt[j] = append(nxt[j], i)
	}

	vis := make([]bool, n)
	cnt := 0
	var hasCycle func(i, parent int) bool
	hasCycle = func(i, parent int) bool {
		if vis[i] {
			return true
		}
		cnt++
		vis[i] = true
		for _, j := range nxt[i] {
			if j != parent && hasCycle(j, i) {
				return true
			}
		}
		return false
	}
	return !hasCycle(0, -1) && cnt == n
}
