package dp

func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	get := func(i int) int {
		if 0 <= i && i < n {
			return nums[i]
		}
		return 1
	}
	for d := 0; d <= n-1; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			sides := get(i-1) * get(j+1)
			for k := i; k <= j; k++ {
				cur := nums[k] * sides
				if i <= k-1 {
					cur += dp[i][k-1]
				}
				if k+1 <= j {
					cur += dp[k+1][j]
				}
				dp[i][j] = max(dp[i][j], cur)
			}
		}
	}
	return dp[0][n-1]
}
