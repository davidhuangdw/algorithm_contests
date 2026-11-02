package main

import (
	"slices"
)

func minCost1(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	dp, pos := make([][]int, n), make([][2]int, 0, n*m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
			pos = append(pos, [2]int{i, j})
		}
	}
	slices.SortFunc(pos, func(a, b [2]int) int {
		return grid[a[0]][a[1]] - grid[b[0]][b[1]]
	})
	dp[n-1][m-1] = 0
	for i := range k + 1 {
		if i > 0 {
			mn := 1 << 60
			for l, r := 0, 0; l < len(pos); l = r {
				for r = l; r < len(pos) && grid[pos[r][0]][pos[r][1]] == grid[pos[l][0]][pos[l][1]]; r++ {
					i, j := pos[r][0], pos[r][1]
					mn = min(dp[i][j], mn)
				}
				for p := l; p < r; p++ {
					i, j := pos[p][0], pos[p][1]
					dp[i][j] = min(dp[i][j], mn)
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			for j := m - 1; j >= 0; j-- {
				cur := &dp[i][j]
				if i+1 < n {
					*cur = min(*cur, dp[i+1][j]+grid[i+1][j])
				}
				if j+1 < m {
					*cur = min(*cur, dp[i][j+1]+grid[i][j+1])
				}
			}
		}
	}
	return dp[0][0]
}
