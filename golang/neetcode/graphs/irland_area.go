package graphs

func maxAreaOfIsland(grid [][]int) int {
	n, m, ans := len(grid), len(grid[0]), 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if !(0 <= i && i < n && 0 <= j && j < m && grid[i][j] == 1) {
			return 0
		}
		grid[i][j] = -1
		return 1 + dfs(i, j-1) + dfs(i, j+1) + dfs(i-1, j) + dfs(i+1, j)
	}
	for i, row := range grid {
		for j := range row {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
