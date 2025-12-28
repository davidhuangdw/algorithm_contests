package graphs

func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	if len(heights) == 0 || len(heights[0]) == 0 {
		return ans
	}

	n, m := len(heights), len(heights[0])
	vis, tmp := make([][]int, n), make([]int, n*m)
	for i, _ := range vis {
		vis[i] = tmp[i*m : i*m+m]
	}

	var dfs func(i, j, flag, h int)
	dfs = func(i, j, flag, h int) {
		if !(0 <= i && i < n && 0 <= j && j < m && vis[i][j]&flag != flag && heights[i][j] >= h) {
			return
		}
		vis[i][j] |= flag
		if vis[i][j] == 3 {
			ans = append(ans, []int{i, j})
		}
		h = heights[i][j]
		dfs(i, j-1, flag, h)
		dfs(i, j+1, flag, h)
		dfs(i-1, j, flag, h)
		dfs(i+1, j, flag, h)

	}
	for i, row := range heights {
		for j, _ := range row {
			flag := 0
			if i == 0 || j == 0 {
				flag |= 1
			}
			if i == n-1 || j == m-1 {
				flag |= 2
			}
			if flag > 0 {
				dfs(i, j, flag, 0)
			}
		}
	}
	return ans
}
