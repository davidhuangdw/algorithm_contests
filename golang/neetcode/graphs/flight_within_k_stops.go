package graphs

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int { // spfa O(k*E)
	nxt := make([][][]int, n)
	for _, e := range flights {
		nxt[e[0]] = append(nxt[e[0]], e)
	}
	dis, INF := make([]int, n), int(1e9)
	for i := range dis {
		dis[i] = INF
	}
	dis[src] = 0
	for t := 0; t < k+1; t++ {
		cur := append([]int(nil), dis...)
		for u, es := range nxt {
			if cur[u] == INF {
				continue
			}
			for _, e := range es {
				v, d := e[1], e[2]
				dis[v] = min(dis[v], cur[u]+d)
			}
		}
	}
	if dis[dst] == INF {
		return -1
	}
	return dis[dst]
}
