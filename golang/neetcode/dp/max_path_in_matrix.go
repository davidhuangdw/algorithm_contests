package dp

func longestIncreasingPath(matrix [][]int) int {
	ans, n, m := 0, len(matrix), len(matrix[0])
	mx := make([][]int, n)
	for i, _ := range mx {
		mx[i] = make([]int, m)
	}
	var dfs func(i, j, pre int) int
	dfs = func(i, j, pre int) int {
		if !(0 <= i && i < n && 0 <= j && j < m && pre < matrix[i][j]) {
			return 0
		}
		if mx[i][j] == 0 {
			v := matrix[i][j]
			mx[i][j] = 1 + max(dfs(i, j-1, v), dfs(i, j+1, v), dfs(i-1, j, v), dfs(i+1, j, v))
		}
		return mx[i][j]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = max(ans, dfs(i, j, -2e9))
		}
	}
	return ans
}
