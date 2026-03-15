package main

func maxActivated(points [][]int) int {
	n := len(points)
	fa, sz := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i], sz[i] = i, 1
	}
	var find func(u int) int
	find = func(u int) int {
		if fa[u] != u {
			fa[u] = find(fa[u])
		}
		return fa[u]
	}
	merge := func(u, v int) {
		u, v = find(u), find(v)
		if sz[u] > sz[v] {
			u, v = v, u
		}
		sz[v] += sz[u]
		fa[u], sz[u] = v, 0
	}

	xp := [2]map[int][]int{}
	for i := range xp {
		xp[i] = make(map[int][]int)
	}

	for u, p := range points {
		for i, v := range p {
			xp[i][v] = append(xp[i][v], u)
		}
	}

	for _, axisPoints := range xp {
		for _, ps := range axisPoints {
			for i := 0; i+1 < len(ps); i++ {
				merge(ps[i], ps[i+1])
			}
		}
	}
	mx, ans := 0, 0
	for _, s := range sz {
		ans = max(ans, mx+s)
		mx = max(mx, s)
	}
	return ans + 1
}
