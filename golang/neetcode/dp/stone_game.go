package dp

func stoneGame1(piles []int) bool {
	if len(piles) == 0 {
		return true
	}
	n := len(piles)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for d := 0; d < n; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			if d == 0 {
				dp[i][j] = piles[i]
			} else {
				dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
			}
		}
	}
	return dp[0][n-1] >= 0
}

func stoneGame(piles []int) bool { // O(n) space
	if len(piles) == 0 {
		return true
	}
	n := len(piles)
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- { // f(i,j) = max(v[i]-f(i+1,j), v[j]-f(i,j-1))
		for j := i; j < n; j++ {
			if i == j {
				dp[j] = piles[j]
			} else {
				dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
			}
		}
	}
	return dp[n-1] >= 0
}
