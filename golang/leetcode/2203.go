package main

import "container/heap"

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	adj, rev := make([][][2]int, n), make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		rev[v] = append(rev[v], [2]int{u, w})
	}

	inf := int(1e18)
	dij := func(st int, g [][][2]int) []int {
		dis := make([]int, n)
		for i := range dis {
			dis[i] = inf
		}
		dis[st] = 0
		q := NewHeap(func(a, b [2]int) bool {
			return a[1] < b[1]
		})
		heap.Push(q, [2]int{st, 0})
		for q.Len() > 0 {
			x := heap.Pop(q).([2]int)
			u, d := x[0], x[1]
			if dis[u] < d {
				continue
			}
			for _, e := range g[u] {
				v, w := e[0], e[1]
				if d+w < dis[v] {
					dis[v] = d + w
					heap.Push(q, [2]int{v, dis[v]})
				}
			}
		}
		return dis
	}
	a, b, c, ans := dij(src1, adj), dij(src2, adj), dij(dest, rev), inf
	for u := range n {
		ans = min(ans, a[u]+b[u]+c[u])
	}
	if ans == inf {
		ans = -1
	}
	return int64(ans)
}
