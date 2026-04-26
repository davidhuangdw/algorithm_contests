package main

func maxCollectedFruits(fruits [][]int) int {
	ans, n := 0, len(fruits)
	for i := range n {
		ans += fruits[i][i]
		fruits[i][i] = 0
	}
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		nxt := dp
		dp = make([]int, n)
		for j := max(i, n-1-i); j < n; j++ {
			cur := &dp[j]
			for nj := j - 1; nj <= min(j+1, n-1); nj++ { // bug: <=j+1
				*cur = max(*cur, nxt[nj])
			}
			*cur += fruits[i][j]
		}
	}
	ans += dp[n-1]
	dp = make([]int, n)
	for j := n - 2; j >= 0; j-- {
		nxt := dp
		dp = make([]int, n)
		for i := max(j, n-1-j); i < n; i++ {
			cur := &dp[i]
			for ni := i - 1; ni <= min(i+1, n-1); ni++ {
				*cur = max(*cur, nxt[ni])
			}
			*cur += fruits[i][j]
		}
	}
	ans += dp[n-1]
	return ans
}
