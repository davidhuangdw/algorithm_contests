package graphs

func islandPerimeter(grid [][]int) int {
	sum := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			sum += 4
			if i-1 >= 0 && grid[i-1][j] == 1 {
				sum -= 2 // decrease overlaps
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				sum -= 2
			}
		}
	}
	return sum
}
