package graphs

func minCostConnectPoints(points [][]int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	dist := func(i, j int) int {
		p := points
		return abs(p[i][0]-p[j][0]) + abs(p[i][1]-p[j][1])
	}

	ans, n := 0, len(points)
	done, dis := make([]bool, n), make([]int, n)
	done[0] = true
	for i := 1; i < n; i++ {
		dis[i] = dist(i, 0)
	}
	for t := 0; t < n-1; t++ {
		mu := -1
		for u := 0; u < n; u++ {
			if !done[u] && (mu < 0 || dis[u] < dis[mu]) {
				mu = u
			}
		}
		ans += dis[mu]
		done[mu] = true
		for v := 0; v < n; v++ {
			if !done[v] {
				dis[v] = min(dis[v], dist(mu, v))
			}
		}
	}
	return ans
}
