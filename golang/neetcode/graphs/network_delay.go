package graphs

func networkDelayTime1(times [][]int, n int, k int) int { // dijkstra O(E+V^2)
	nxt := make([][][]int, n+1)
	for _, ti := range times {
		u, v, d := ti[0], ti[1], ti[2]
		nxt[u] = append(nxt[u], []int{v, d})
	}
	dis, mx, INF := make([]int, n+1), 0, int(1e9)
	for i := 1; i <= n; i++ {
		dis[i] = INF
	}
	dis[k] = 0
	for _t := 0; _t < n; _t++ {
		mn, u := INF, -1
		for i := 1; i <= n; i++ {
			if dis[i] >= 0 && dis[i] < mn {
				mn, u = dis[i], i
			}
		}
		if mn == INF {
			return -1
		}
		mx = max(mx, dis[u])
		dis[u] = -1 // done
		for _, p := range nxt[u] {
			v, d := p[0], p[1]
			if dis[v] >= 0 {
				dis[v] = min(dis[v], mn+d)
			}
		}
	}
	return mx
}

func networkDelayTime(times [][]int, n int, k int) int { // spfa, O(V*E) worst
	nxt := make([][][]int, n+1)
	for _, ti := range times {
		u, v, d := ti[0], ti[1], ti[2]
		nxt[u] = append(nxt[u], []int{v, d})
	}

	INF := int(1e9)
	dis, qued := make([]int, n+1), make([]bool, n+1)
	for i := 1; i <= n; i++ {
		dis[i] = INF
	}
	dis[k], qued[k] = 0, true
	que := []int{k}
	for len(que) > 0 {
		u := que[0]
		que, qued[u] = que[1:], false
		for _, p := range nxt[u] {
			v, d := p[0], p[1]
			if dis[u]+d < dis[v] {
				dis[v] = dis[u] + d
				if !qued[v] {
					que = append(que, v)
					qued[v] = true
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dis[i] == INF {
			return -1
		}
		ans = max(ans, dis[i])
	}
	return ans
}
