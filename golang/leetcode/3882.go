package main

func minCost(grid [][]int) int {
	n, m, mx := len(grid), len(grid[0]), 1024
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, m)
		for j := range dp[i] {
			dp[i][j] = make([]bool, mx)
			if i == 0 && j == 0 {
				dp[i][j][grid[i][j]] = true
				continue
			}
			for v := range dp[i][j] {
				tar := v ^ grid[i][j]
				dp[i][j][v] = (i-1 >= 0 && dp[i-1][j][tar]) || (j-1 >= 0 && dp[i][j-1][tar])
			}
		}
	}

	for v, ok := range dp[n-1][m-1] {
		if ok {
			return v
		}
	}

	return -1
}
