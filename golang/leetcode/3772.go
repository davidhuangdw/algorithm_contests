package main

func maxSubgraphScore(n int, edges [][]int, good []int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}
	for i, v := range good {
		if v == 0 {
			good[i] = -1
		}
	}

	sub := make([]int, n)
	var getSub func(u, p int) int
	getSub = func(u, p int) int {
		sub[u] += good[u]
		for _, v := range adj[u] {
			if v != p {
				sub[u] += max(0, getSub(v, u))
			}
		}
		return sub[u]
	}
	getSub(0, -1)

	ans := make([]int, n)
	var getFull func(u, p, vp int)
	getFull = func(u, p, vp int) {
		ans[u] = max(0, vp-max(0, sub[u])) + sub[u]
		for _, v := range adj[u] {
			if v != p {
				getFull(v, u, ans[u])
			}
		}
	}
	getFull(0, -1, 0)
	return ans
}
