package graphs

import "container/heap"

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	que := NewHeap(func(a, b []int) bool {
		return a[0] < b[0]
	})
	heap.Push(que, []int{0, 0, 0})
	for que.Len() > 0 {
		v := heap.Pop(que).([]int)
		d, i, j := v[0], v[1], v[2]
		if res[i][j] >= 0 {
			continue
		}
		res[i][j] = d
		di, dj := 0, 1
		for t := 0; t < 4; t++ {
			ni, nj := i+di, j+dj
			if 0 <= ni && ni < n && 0 <= nj && nj < m && res[ni][nj] < 0 {
				nd := heights[ni][nj] - heights[i][j]
				nd = max(nd, -nd)
				heap.Push(que, []int{max(d, nd), ni, nj})
			}
			di, dj = dj, -di
		}
	}
	return res[n-1][m-1]
}
