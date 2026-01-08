package dp

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 2e9
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[n]
}
