package main

func numMusicPlaylists(n int, goal int, k int) int {
	M := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[0] = 1
	for l := 1; l <= goal; l++ {
		pre := dp[0]
		dp[0] = 0
		for t := 1; t <= min(n, l); t++ { // exactly t types
			dp[t], pre = (pre*(n-(t-1))+dp[t]*max(0, t-k))%M, dp[t]
		}
	}
	return dp[n]
}
