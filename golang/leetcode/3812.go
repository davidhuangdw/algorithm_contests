package main

import "slices"

func minimumFlips(n int, edges [][]int, start string, target string) []int {
	adj := make([][]int, n)
	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], i)
		adj[e[1]] = append(adj[e[1]], i)
	}
	var ans []int
	var dfs func(u, p int) bool
	dfs = func(u, p int) bool {
		f := 0
		if start[u] != target[u] {
			f = 1
		}
		for _, ei := range adj[u] {
			e := edges[ei]
			v := u ^ e[0] ^ e[1]
			if v != p {
				if dfs(v, u) {
					f ^= 1
					ans = append(ans, ei)
				}
			}
		}
		return f == 1
	}
	if dfs(0, -1) {
		ans = []int{-1}
	}
	slices.Sort(ans)
	return ans
}
