package main

import (
	"slices"
)

func maxIncreasingCells1(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	ids := make([][]int, n)
	for i := range ids {
		ids[i] = make([]int, m)
		for j := range ids[i] {
			ids[i][j] = i*m + j
		}
	}

	adj := make([][]int, n*m)
	// row
	for i := range mat {
		ind := make(map[int][]int)
		var vs []int
		for j, v := range mat[i] {
			ind[v] = append(ind[v], j)
		}
		for v := range ind {
			vs = append(vs, v)
		}
		slices.Sort(vs)
		for k, pre := len(vs)-1, -1; k >= 0; k-- {
			js := ind[vs[k]]
			if pre >= 0 {
				for _, j := range js {
					id := ids[i][j]
					adj[id] = append(adj[id], pre)
				}
			}
			// pre id
			if len(js) == 1 {
				pre = ids[i][js[0]]
			} else {
				pre = len(adj) // new id
				adj = append(adj, nil)
				for _, j := range js {
					adj[pre] = append(adj[pre], ids[i][j])
				}
			}
		}
	}
	// col
	for j := 0; j < m; j++ {
		ind := make(map[int][]int)
		var vs []int
		for i := 0; i < n; i++ {
			v := mat[i][j]
			ind[v] = append(ind[v], i)
		}
		for v := range ind {
			vs = append(vs, v)
		}
		slices.Sort(vs)
		for k, pre := len(vs)-1, -1; k >= 0; k-- {
			is := ind[vs[k]]
			if pre >= 0 {
				for _, i := range is {
					id := ids[i][j]
					adj[id] = append(adj[id], pre)
				}
			}
			// pre id
			if len(is) == 1 {
				pre = ids[is[0]][j]
			} else {
				pre = len(adj) // new id
				adj = append(adj, nil)
				for _, i := range is {
					adj[pre] = append(adj[pre], ids[i][j])
				}
			}
		}
	}

	// topo
	ans, dp := 0, make([]int, len(adj))
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(u int) int {
		if dp[u] < 0 {
			dp[u] = 0
			for _, v := range adj[u] {
				dp[u] = max(dp[u], dfs(v))
			}
			if u < n*m {
				dp[u]++
			}
		}
		return dp[u]
	}
	for i := 0; i < len(adj); i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

func maxIncreasingCells(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	vs := make([][3]int, 0, n*m)
	for i := range mat {
		for j, v := range mat[i] {
			vs = append(vs, [3]int{v, i, j})
		}
	}
	slices.SortFunc(vs, func(a, b [3]int) int {
		return a[0] - b[0]
	})
	row, col := make([]int, n), make([]int, m)
	for l, r := 0, 0; l < len(vs); l = r {
		for r = l + 1; r < len(vs) && vs[r-1][0] == vs[r][0]; r++ {
		}
		res := make([]int, 0, r-l)
		for _, p := range vs[l:r] {
			i, j := p[1], p[2]
			res = append(res, 1+max(row[i], col[j]))
		}
		for t, p := range vs[l:r] {
			i, j := p[1], p[2]
			row[i] = max(row[i], res[t])
			col[j] = max(col[j], res[t])
		}
	}
	return slices.Max(row)
}
