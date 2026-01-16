package graphs

func canTraverseAllPairs(nums []int) bool {
	mx := 1
	for _, v := range nums {
		if v == 1 { // bug: forgot
			return false
		}
		mx = max(mx, v)
	}
	comp := make([]int, mx+1) // its least prime factor
	for p := 2; p <= mx; p++ {
		if comp[p] != 0 {
			continue
		}
		comp[p] = p
		for v := p * p; v <= mx; v += p {
			if comp[v] == 0 {
				comp[v] = p
			}
		}
	}

	conn := make(map[int]map[int]bool)
	get := func(p int) map[int]bool {
		if _, ok := conn[p]; !ok {
			conn[p] = make(map[int]bool)
		}
		return conn[p]
	}
	add := func(p, q int) {
		get(p)[q] = true
		get(q)[p] = true
	}

	for _, x := range nums {
		for comp[x] > 0 {
			p := comp[x]
			get(p)
			x /= p
			for comp[x] == p {
				x /= p
			}
			if x == 1 {
				break
			}
			add(p, comp[x])
		}
	}
	vis := make(map[int]bool)
	var dfs func(p, last int)
	dfs = func(p, last int) {
		if vis[p] {
			return
		}
		vis[p] = true
		for q := range conn[p] {
			if q != last {
				dfs(q, p)
			}
		}
	}
	dfs(comp[nums[0]], -1)
	return len(vis) == len(conn)
}
