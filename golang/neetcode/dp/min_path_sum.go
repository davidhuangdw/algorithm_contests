package dp

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([]int, m)
	for i, row := range grid {
		for j, v := range row {
			var add int
			switch {
			case j == 0:
				add = dp[j]
			case i == 0:
				add = dp[j-1]
			default:
				add = min(dp[j-1], dp[j])
			}
			dp[j] = v + add
		}
	}
	return dp[m-1]
}
