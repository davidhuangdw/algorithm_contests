package backtrace

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)

	row, atK := make([]byte, n), make([]string, 0, n)
	for i := 0; i < n; i++ {
		row[i] = '.'
	}
	for i := 0; i < n; i++ {
		row[i] = 'Q'
		atK = append(atK, string(row))
		row[i] = '.'
	}

	mj, add, minus := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var dfs func(i int, pre []string)
	dfs = func(i int, pre []string) {
		if i >= n {
			ans = append(ans, append([]string(nil), pre...)) // bug: should clone pre
			return
		}
		for j := 0; j < n; j++ {
			if mj[j] || add[i+j] || minus[i-j] {
				continue
			}
			mj[j], add[i+j], minus[i-j] = true, true, true
			dfs(i+1, append(pre, atK[j]))
			mj[j], add[i+j], minus[i-j] = false, false, false
		}
	}
	dfs(0, make([]string, 0))
	return ans
}
