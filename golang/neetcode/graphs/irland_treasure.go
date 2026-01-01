package graphs

func islandsAndTreasure(grid [][]int) {
	OK := (1 << 31) - 1
	que, n, m := make([][]int, 0), len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				que = append(que, []int{i, j, 0})
			}
		}
	}
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for di, dj, k := 0, 1, 0; k < 4; di, dj, k = dj, -di, k+1 {
			i, j := p[0]+di, p[1]+dj
			if 0 <= i && i < n && 0 <= j && j < m && grid[i][j] == OK {
				dis := p[2] + 1
				grid[i][j] = dis
				que = append(que, []int{i, j, dis})
			}
		}
	}
}
