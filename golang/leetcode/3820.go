package main

func specialNodes(n int, edges [][]int, x int, y int, z int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	getDis := func(r int) map[int]int {
		dis := make(map[int]int)
		var dfs func(u, p, d int)
		dfs = func(u, p, d int) {
			dis[u] = d
			for _, v := range adj[u] {
				if v != p {
					dfs(v, u, d+1)
				}
			}
		}
		dfs(r, -1, 0)
		return dis
	}
	dx := getDis(x)
	dy := getDis(y)
	dz := getDis(z)
	ans := 0
	for u := 0; u < n; u++ {
		a, b, c := dx[u], dy[u], dz[u]
		if a < b {
			a, b = b, a
		}
		if a < c {
			a, c = c, a
		}
		if a*a == b*b+c*c {
			ans++
		}
	}
	return ans
}
