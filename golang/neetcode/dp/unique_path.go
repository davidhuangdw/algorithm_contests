package dp

func uniquePaths1(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for ; m > 0; m-- {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}

func uniquePaths(m int, n int) int { // C(m+n-2, n-1)
	m, n, ans := min(m, n), max(m, n), 1
	for i, j := m, 1; i <= m+n-2; i, j = i+1, j+1 {
		ans = ans * i / j // okay to '/j' since already '*i' j times
	}
	return ans
}
