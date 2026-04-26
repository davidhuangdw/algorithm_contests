package main

func maxScore(grid [][]int) int {
	rows := make(map[int]int)
	for i, row := range grid {
		r := 1 << i
		for _, v := range row {
			rows[v] |= r
		}
	}
	B := 1 << len(grid)
	dp := make([]int, B)
	for v, msk := range rows {
		for b := B - 1; b >= 0; b-- {
			for rs, lb := msk-(msk&b), 0; rs > 0; rs -= lb {
				lb = rs & -rs
				dp[b|lb] = max(dp[b|lb], dp[b]+v)
			}
		}
	}
	return dp[B-1]
}
