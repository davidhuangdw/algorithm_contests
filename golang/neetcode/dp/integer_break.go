package dp

import "math"

func integerBreak1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, n-1); j++ {
			dp[i] = max(dp[i], dp[i-j]*j)
		}
	}
	return dp[n]
}

func integerBreak2(n int) int {
	dp := make([][2]int, n+1)
	dp[0] = [2]int{1, 1}
	for i := 0; i < n; i++ {
		prd, st := dp[i][0], dp[i][1]
		if i >= prd {
			prd, st = i, i
		}
		for j := min(n-i, st); j >= 1; j-- { // dp with less try: till st
			if prd*j >= dp[i+j][0] {
				dp[i+j] = [2]int{prd * j, j}
			}
		}
	}
	return dp[n][0]
}

func integerBreak(n int) int { // by math
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 1 {
		b = 4 // as 2*2
		a -= 1
	}
	if b == 0 {
		b = 1
	}
	return int(math.Pow(3, float64(a))) * b
}
