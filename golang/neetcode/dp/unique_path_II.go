package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid[0])
	dp := make([]int, m)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for j, v := range row {
			if v == 1 {
				dp[j] = 0
			} else if j-1 >= 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[m-1]
}
