package _2065

import "container/heap"

const MAX = 1e9

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, t})
		adj[v] = append(adj[v], edge{u, t})
	}
	dis := minDis(adj, 0)

	mx, sum, time := 0, 0, 0

	vis := map[int]bool{}

	var dfs func(u int)
	dfs = func(u int) {
		if !vis[u] {
			sum += values[u]
			vis[u] = true
			defer func() {
				sum -= values[u]
				vis[u] = false
			}()
		}

		if u == 0 && sum > mx {
			mx = sum
		}
		for _, e := range adj[u] {
			v, t := e.v, e.dis
			if time+t+dis[v] <= maxTime {
				time += t
				dfs(v)
				time -= t
			}
		}
	}
	dfs(0)
	return mx
}

type edge struct {
	v, dis int
}

type end struct {
	u, dis int
}

type Que []end

func (q Que) Less(i, j int) bool { return q[i].dis < q[j].dis }
func (q Que) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q Que) Len() int           { return len(q) }

func (q *Que) Push(x interface{}) { *q = append(*q, x.(end)) }
func (q *Que) Pop() interface{} {
	x := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return x
}

func minDis(adj [][]edge, st int) []int { //dijkstra
	n := len(adj)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = MAX
	}
	var que Que
	dis[st] = 0 // bug: forgot
	heap.Push(&que, end{st, dis[st]})
	for len(que) > 0 {
		x := heap.Pop(&que).(end)
		if dis[x.u] < x.dis {
			continue
		}
		for _, e := range adj[x.u] {
			if x.dis+e.dis < dis[e.v] {
				dis[e.v] = x.dis + e.dis
				heap.Push(&que, end{e.v, dis[e.v]})
			}
		}
	}
	return dis
}

//func minDis(adj [][]edge, st int) []int { //spfa
//	n := len(adj)
//	dis := make([]int, n)
//	for i := range dis{
//		dis[i] = MAX
//	}
//
//	var que []int
//	queued := map[int]bool{}
//
//	dis[st] = 0
//	que = append(que, st)
//	queued[st] = true
//
//	for len(que) > 0 {
//		u := que[0]
//		que = que[1:]
//		queued[u] = false 	// bug: forgot
//
//		for _, e := range adj[u] {
//			v, delta := e.v, e.dis
//			if dis[u] + delta < dis[v] {
//				dis[v] = dis[u] + delta
//				if !queued[v] {
//					que = append(que, v)
//					queued[v] = true
//				}
//			}
//		}
//	}
//	return dis
//}
