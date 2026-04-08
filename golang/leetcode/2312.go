package main

func sellingWood(m int, n int, prices [][]int) int64 {
	pr := make(map[[2]int]int)
	for _, p := range prices {
		sz := [2]int{p[0], p[1]}
		pr[sz] = max(pr[sz], p[2])
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			v := &dp[i][j]
			*v = pr[[2]int{i, j}]
			for l := 1; l <= j/2; l++ {
				*v = max(*v, dp[i][l]+dp[i][j-l])
			}
			for t := 1; t <= i/2; t++ {
				*v = max(*v, dp[t][j]+dp[i-t][j])
			}
		}
	}
	return int64(dp[m][n])
}
