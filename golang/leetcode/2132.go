package main

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	h, w, n, m := stampHeight, stampWidth, len(grid), len(grid[0])
	sm, diff := make([][]int, n+1), make([][]int, n+1)
	for i := range sm {
		sm[i] = make([]int, m+1)
		diff[i] = make([]int, m+1)
		for j := range sm[i] {
			if i > 0 && j > 0 {
				sm[i][j] = sm[i][j-1] + sm[i-1][j] - sm[i-1][j-1] + grid[i-1][j-1]
			}
			if i >= h && j >= w && sm[i][j]-sm[i-h][j]-sm[i][j-w]+sm[i-h][j-w] == 0 {
				diff[i][j]++
				diff[i-h][j]--
				diff[i][j-w]--
				diff[i-h][j-w]++
			}
		}
	}
	for i := range n {
		for j := range m {
			sm[i+1][j+1] = sm[i][j+1] + sm[i+1][j] - sm[i][j] + diff[i][j]
			if sm[i+1][j+1] == 0 && grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
