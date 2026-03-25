package main

func interactionCosts(n int, edges [][]int, group []int) int64 {
	ans, adj := 0, make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}

	var gr, total, cnt int
	var dfs func(u, p int)
	dfs = func(u, p int) {
		pc := cnt
		total += pc
		if group[u] == gr {
			ans += total
			cnt++
		}
		for _, v := range adj[u] {
			if v != p {
				dfs(v, u)
			}
		}
		total += (cnt - pc) - pc
		return
	}

	grp := make(map[int]bool)
	for _, g := range group {
		grp[g] = true
	}
	for g := range grp {
		gr, total, cnt = g, 0, 0
		dfs(0, -1)
	}
	return int64(ans)
}
