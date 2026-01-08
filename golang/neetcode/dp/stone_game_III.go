package dp

func stoneGameIII(stones []int) string {
	n := len(stones)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = -1e9
		sum := 0
		for j := i; j < min(n, i+3); j++ {
			sum += stones[j]
			dp[i] = max(dp[i], sum-dp[j+1])
		}
	}

	switch {
	case dp[0] == 0:
		return "Tie"
	case dp[0] > 0:
		return "Alice"
	default:
		return "Bob"
	}
}
