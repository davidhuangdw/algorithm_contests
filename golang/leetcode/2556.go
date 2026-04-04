package main

func isPossibleToCutPath(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	var run func(i, j int) bool
	run = func(i, j int) bool {
		if i == n-1 && j == m-1 {
			return true
		}
		grid[i][j] = 0
		return (i+1 < n && grid[i+1][j] == 1 && run(i+1, j)) ||
			(j+1 < m && grid[i][j+1] == 1 && run(i, j+1))
	}
	return !(run(0, 0) && run(0, 0))
}
