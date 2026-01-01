package graphs

func findOrder(n int, prerequisites [][]int) []int {
	nxt, vis := make(map[int][]int), make([]int, n)
	var ans []int
	for _, p := range prerequisites {
		i, j := p[0], p[1]
		nxt[i] = append(nxt[i], j)
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if vis[i] != 0 {
			return vis[i] == 2
		}
		vis[i] = 1
		for _, j := range nxt[i] {
			if !dfs(j) {
				return false
			}
		}
		ans = append(ans, i)
		vis[i] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if !dfs(i) {
			return nil
		}
	}
	return ans
}
