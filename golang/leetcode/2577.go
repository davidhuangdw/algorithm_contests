package main

import "container/heap"

func minimumTime(grid [][]int) int {
	n, m, inf := len(grid), len(grid[0]), int(1e9)
	dis, vis := make([][]int, n), make([][]bool, n)
	for i := range dis {
		dis[i], vis[i] = make([]int, m), make([]bool, m)
		for j := range dis[i] {
			dis[i][j] = inf
		}
	}
	vis[0][0] = true
	q := NewHeap(func(a, b [3]int) bool {
		return a[0] < b[0]
	})
	if grid[0][1] <= 1 {
		heap.Push(q, [3]int{1, 0, 1})
	}
	if grid[1][0] <= 1 {
		heap.Push(q, [3]int{1, 1, 0})
	}
	for q.Len() > 0 {
		v := heap.Pop(q).([3]int)
		sd, si, sj := v[0], v[1], v[2]
		if vis[si][sj] {
			continue
		}
		vis[si][sj] = true
		dis[si][sj] = sd
		if si == n-1 && sj == m-1 {
			return sd
		}
		sd++
		for t, di, dj := 0, 0, 1; t < 4; t++ {
			i, j := si+di, sj+dj
			if 0 <= i && i < n && 0 <= j && j < m && !vis[i][j] {
				d := sd + max(0, (grid[i][j]-sd+1)/2*2)
				heap.Push(q, [3]int{d, i, j})
			}
			di, dj = dj, -di
		}
	}
	return -1
}
