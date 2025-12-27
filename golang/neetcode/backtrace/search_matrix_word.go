package backtrace

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	wd, n, m := []byte(word), len(board), len(board[0])

	var dfs func(i, x, y int) bool
	dfs = func(i, x, y int) bool {
		if i >= len(wd) {
			return true
		}
		if !(0 <= x && x < n && 0 <= y && y < m && board[x][y] == wd[i]) {
			return false
		}

		origin := board[x][y]
		board[x][y] = '#'
		i++
		if dfs(i, x-1, y) || dfs(i, x+1, y) || dfs(i, x, y-1) || dfs(i, x, y+1) {
			return true
		}
		board[x][y] = origin
		return false
	}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if dfs(0, x, y) {
				return true
			}
		}
	}
	return false
}
