package graphs

func orangesRotting(grid [][]int) int {
	n, m, cnt, mx := len(grid), len(grid[0]), 0, 0
	var que [][]int
	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				que = append(que, []int{i, j, 0})
			}
			if v == 1 {
				cnt++
			}
		}
	}
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		di, dj := 0, 1
		for k := 0; k < 4; k++ {
			i, j := p[0]+di, p[1]+dj
			if 0 <= i && i < n && 0 <= j && j < m && grid[i][j] == 1 {
				cnt--
				grid[i][j] = 2
				dis := p[2] + 1
				mx = max(mx, dis)
				que = append(que, []int{i, j, dis})
			}
			di, dj = dj, -di
		}
	}

	if cnt > 0 {
		return -1
	}
	return mx
}
