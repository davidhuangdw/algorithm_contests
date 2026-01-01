package graphs

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !(0 <= i && i < n && 0 <= j && j < m && board[i][j] == 'O') {
			return
		}
		board[i][j] = '#'
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs(0, j)
		dfs(n-1, j)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '#':
				board[i][j] = 'O'
			}
		}
	}
}
