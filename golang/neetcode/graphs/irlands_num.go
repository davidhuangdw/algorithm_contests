package graphs

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m, ans := len(grid), len(grid[0]), 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !(0 <= i && i < n && 0 <= j && j < m && grid[i][j] == '1') {
			return
		}
		grid[i][j] = '#'
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	for i, row := range grid {
		for j, ch := range row {
			if ch == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	// maybe restore '#' back to '1' at last
	return ans
}
