package graphs

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 { // O(V+E) dfs
	nodes, ans := make(map[string]map[string]float64), make([]float64, len(queries))
	nodeQueries := make(map[string][]int)
	for i, q := range queries {
		u, v := q[0], q[1]
		nodeQueries[u] = append(nodeQueries[u], i)
		nodeQueries[v] = append(nodeQueries[v], i)
		ans[i] = -1
	}
	getNode := func(u string) map[string]float64 {
		if _, ok := nodes[u]; !ok {
			nodes[u] = make(map[string]float64)
		}
		return nodes[u]
	}
	adj := make(map[string][]string)
	for i, eq := range equations {
		u, v := eq[0], eq[1]
		d := values[i]
		getNode(u)[v] = 1 / d
		getNode(v)[u] = d
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	vis := make(map[string]bool)
	var relative map[string]float64
	var dfs func(u string, ux float64)
	dfs = func(u string, ux float64) {
		if vis[u] {
			return
		}
		vis[u] = true
		relative[u] = ux
		// answer
		for _, i := range nodeQueries[u] {
			q := queries[i]
			a, b := relative[q[0]], relative[q[1]]
			if a > 0 && b > 0 {
				ans[i] = a / b
			}
		}
		// expand
		for _, v := range adj[u] {
			dfs(v, ux*getNode(u)[v])
		}
	}
	for u := range nodes {
		if vis[u] {
			continue
		}
		relative = make(map[string]float64)
		dfs(u, 1)
	}
	return ans
}
