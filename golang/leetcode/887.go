package main

func superEggDrop1(k int, n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	for i := 2; i <= k; i++ {
		ndp := make([]int, n+1)
		ndp[1] = 1
		for j := 2; j <= n; j++ {
			l, r := 1, j
			for l <= r {
				md := (l + r) >> 1
				if ndp[j-md] <= dp[md-1] {
					r = md - 1
				} else {
					l = md + 1
				}
			}
			ndp[j] = 1 + max(dp[l-1], ndp[j-l])
		}
		dp = ndp
	}
	return dp[n]
}

func superEggDrop(k int, n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = 1
	}
	for t := 2; ; t++ {
		for i := k; i >= 1; i-- {
			dp[i] += 1 + dp[i-1]
		}
		if dp[k] >= n {
			return t
		}
	}
}
