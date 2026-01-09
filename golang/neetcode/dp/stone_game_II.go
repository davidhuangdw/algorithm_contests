package dp

func stoneGameII(piles []int) int {
	// f(k, i) = max{sum(i:j)-f(max(j-i,k), j)}, j in [i+1, i+k*2]

	if len(piles) == 0 {
		return 0
	}
	n := len(piles)
	dp, total := make([][]int, n+1), 0
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range piles {
		total += v
	}
	for i := n - 1; i >= 0; i-- {
		for k := 1; k <= n; k++ {
			sum := 0
			dp[i][k] = -1e9
			for j := i + 1; j <= min(n, i+2*k); j++ {
				sum += piles[j-1]
				dp[i][k] = max(dp[i][k], sum-dp[j][max(j-i, k)])
			}
		}
	}

	return (total + dp[0][1]) / 2
}
