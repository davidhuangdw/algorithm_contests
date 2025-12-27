package backtrace

func generateParenthesis(n int) (ans []string) {
	var dfs func(pre, half int, cur []byte)
	dfs = func(pre, half int, cur []byte) {
		if pre >= n && half == 0 {
			ans = append(ans, string(cur))
			return
		}
		if pre < n {
			dfs(pre+1, half+1, append(cur, '('))
		}
		if half > 0 {
			dfs(pre, half-1, append(cur, ')'))
		}
	}

	dfs(0, 0, nil)
	return ans
}
