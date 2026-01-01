package dp

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := m - 1; i >= 0; i-- {
		pre := 1
		for j := n - 1; j >= 0; j-- {
			tmp := dp[j]
			if s[i] == t[j] {
				dp[j] += pre
			}
			pre = tmp
		}
	}
	return dp[0]
}
