package dp

func isMatch1(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([]bool, n+1)
	equal := func(i, j int) bool {
		return p[i] == '.' || p[i] == s[j]
	}
	for i := m; i >= 0; i-- { // pi
		var pre bool
		for j := n; j >= 0; j-- { // sj
			cur := false
			switch {
			case i == m && j == n:
				cur = true
			case i == m:
				cur = false
			case p[i] == '*':
				cur = dp[j]
			case i+1 < m && p[i+1] == '*':
				cur = dp[j] || (j < n && equal(i, j) && dp[j+1])
			case j == n:
				cur = false
			default:
				cur = equal(i, j) && pre
			}
			pre, dp[j] = dp[j], cur
		}
	}
	return dp[0]
}

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([]bool, m+1)
	equal := func(i, j int) bool {
		return i < n && j < m && (p[j] == '.' || s[i] == p[j])
	}
	for i := n; i >= 0; i-- {
		var pre bool
		for j := m; j >= 0; j-- {
			var v bool
			switch {
			case j == m: // s[i:], ''
				v = i == n
			case p[j] == '*': // s[i:], '*...'
				v = dp[j+1]
			case j+1 < m && p[j+1] == '*': // s[i:], 'x*...'
				v = dp[j+1] || (equal(i, j) && dp[j])
			default: // s[i:], 'x...'
				v = equal(i, j) && pre
			}
			dp[j], pre = v, dp[j]
		}
	}
	return dp[0]
}
