package main

func minDistance(word1 string, word2 string) int {
	m := len(word2)
	dp, pre := make([]int, m+1), 0
	for i := range dp {
		dp[i] = i
	}
	for i, a := range word1 {
		dp[0], pre = i+1, dp[0]
		for j, b := range word2 {
			if a == b {
				pre, dp[j+1] = dp[j+1], pre
			} else {
				pre, dp[j+1] = dp[j+1], min(pre, dp[j+1], dp[j])+1
			}
		}
	}
	return dp[m]
}
