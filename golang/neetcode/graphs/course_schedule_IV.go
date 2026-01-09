package graphs

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, n)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}
	cache := make([]map[int]bool, n)

	var vis map[int]bool
	var dfs func(u int)
	dfs = func(u int) {
		if vis[u] {
			return
		}
		vis[u] = true
		for _, v := range adj[u] {
			dfs(v)
		}
	}
	collect := func(u int) map[int]bool {
		if cache[u] == nil {
			vis = make(map[int]bool)
			dfs(u)
			cache[u] = vis
		}
		return cache[u]
	}
	ans := make([]bool, len(queries))
	for i, p := range queries {
		u, v := p[0], p[1]
		ans[i] = u != v && collect(u)[v]
	}
	return ans
}
