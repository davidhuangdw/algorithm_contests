package graphs

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	topoPos := func(cond [][]int) ([]int, bool) {
		// buid graph
		nxt := make([][]int, k+1)
		for _, e := range cond {
			v, u := e[0], e[1] // high->low
			nxt[u] = append(nxt[u], v)
		}

		// build topo
		vis := make([]int, k+1) //0-not visit, 1-visiting, 2-done
		var topo []int
		var dfs func(u int) bool
		dfs = func(u int) bool {
			if vis[u] != 0 {
				return vis[u] != 1
			}
			vis[u] = 1
			for _, v := range nxt[u] {
				if !dfs(v) {
					return false
				}
			}
			topo = append(topo, u)
			vis[u] = 2
			return true
		}
		for u := 1; u <= k; u++ {
			if !dfs(u) {
				return nil, false
			}
		}

		// topo to pos
		pos := make([]int, k+1)
		for i, u := range topo {
			pos[u] = i
		}
		return pos, true
	}

	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}
	row, ok := topoPos(rowConditions)
	if !ok {
		return nil
	}

	col, ok := topoPos(colConditions)
	if !ok {
		return nil
	}
	for v := 1; v <= k; v++ {
		ans[row[v]][col[v]] = v
	}
	return ans
}
