package main

func maxSizeSlices(slices []int) int {
	l := len(slices) / 3
	dp := func(a []int) int {
		n := len(a)
		dp := make([]int, n+1)
		for t := 1; t <= l; t++ {
			pre := dp
			dp = make([]int, n+1)
			for i := 2*t - 1; i <= n; i++ {
				use := a[i-1]
				if i-2 >= 0 {
					use += pre[i-2]
				}
				dp[i] = max(dp[i-1], use)
			}
		}
		return dp[n]
	}
	return max(dp(slices[1:]), dp(slices[:len(slices)-1]))
}
