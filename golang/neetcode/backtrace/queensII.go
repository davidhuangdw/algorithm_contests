package backtrace

func totalNQueens(n int) (ans int) {
	mj, add, minus := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			if mj[j] || add[i+j] || minus[i-j] {
				continue
			}
			mj[j] = true
			add[i+j] = true
			minus[i-j] = true
			dfs(i + 1)
			mj[j] = false
			add[i+j] = false
			minus[i-j] = false
		}
	}
	dfs(0)
	return ans
}
