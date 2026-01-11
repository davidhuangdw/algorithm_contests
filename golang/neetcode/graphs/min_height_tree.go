package graphs

func findMinHeightTrees1(n int, edges [][]int) (ans []int) { // by reroot
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	ans = []int{0}
	mx, mxLeaf, ready := 0, 0, false
	var dfs func(u, par int, path []int)
	dfs = func(u, par int, path []int) { // search farthest leaf
		if len(adj[u]) <= 1 && len(path) > mx {
			mx = len(path)
			mxLeaf = u
			if ready {
				l := len(path)
				if l%2 == 0 {
					ans = []int{path[l/2-1], path[l/2]}
				} else {
					ans = []int{path[(l-1)/2]}
				}
			}
		}
		for _, v := range adj[u] {
			if v != par {
				dfs(v, u, append(path, v))
			}
		}
	}
	dfs(0, -1, []int{0})
	ready, mx = true, 0
	dfs(mxLeaf, -1, []int{mxLeaf})
	return ans
}

func findMinHeightTrees(n int, edges [][]int) (ans []int) { // by degree layered-bfs
	deg, adj := make([]int, n), make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	qued := make([]bool, n)
	var leaf []int
	for u := range adj {
		deg[u] = len(adj[u])
		if deg[u] <= 1 {
			leaf = append(leaf, u)
			qued[u] = true
		}
	}
	remain := n

	for remain > 2 {
		var nxtLeaf []int
		remain -= len(leaf)
		for _, u := range leaf {
			for _, v := range adj[u] {
				if qued[v] {
					continue
				}
				deg[v]--
				if deg[v] == 1 {
					qued[v] = true
					nxtLeaf = append(nxtLeaf, v)
				}
			}
		}
		leaf = nxtLeaf
	}
	return leaf
}
