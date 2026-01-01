package dp

func minDistance(a string, b string) int {
	n, m := len(a), len(b)
	dp := make([]int, m+1)
	for i := n; i >= 0; i-- {
		var pre, old int
		for j := m; j >= 0; j-- {
			if i == n || j == m {
				pre = dp[j]
				dp[j] = (n - i) + (m - j)
				continue
			}
			old = dp[j]
			dp[j] = 1 + min(dp[j], dp[j+1], pre)
			if a[i] == b[j] {
				dp[j] = pre
			}
			pre = old
		}
	}
	return dp[0]
}
