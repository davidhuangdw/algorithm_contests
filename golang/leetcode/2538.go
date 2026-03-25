package main

func maxOutput1(n int, edges [][]int, price []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}

	subMax, subMin := make([]int, n), make([]int, n)
	var buildSub func(u, p int)
	buildSub = func(u, p int) {
		mx, mn := 0, 0 // not use child
		for _, v := range adj[u] {
			if v != p {
				buildSub(v, u)
				mx = max(mx, subMax[v])
				mn = min(mn, subMin[v])
			}
		}
		subMax[u], subMin[u] = price[u]+mx, price[u]+mn
	}
	buildSub(0, -1)

	ans := 0
	var dfs func(u, p, upMax, upMin int)
	dfs = func(u, p, upMax, upMin int) {
		ans = max(ans, max(upMax+price[u], subMax[u])-min(upMin+price[u], subMin[u]))

		preMx, sufMx := make([]int, len(adj[u])), make([]int, len(adj[u]))
		preMn, sufMn := make([]int, len(adj[u])), make([]int, len(adj[u]))
		pMx, pMn := 0, 0
		for i, v := range adj[u] {
			preMx[i], preMn[i] = pMx, pMn
			if v != p {
				pMx = max(pMx, subMax[v])
				pMn = max(pMn, subMin[v])
			}
		}
		pMx, pMn = 0, 0
		for i := len(adj[u]) - 1; i >= 0; i-- {
			v := adj[u][i]
			sufMx[i], sufMn[i] = pMx, pMn
			if v != p {
				pMx = max(pMx, subMax[v])
				pMn = max(pMn, subMin[v])
			}
		}

		for i, v := range adj[u] {
			if v != p {
				dfs(v, u, price[u]+max(upMax, preMx[i], sufMx[i]), price[u]+min(upMin, preMn[i], sufMn[i]))
			}
		}
	}
	dfs(0, -1, 0, 0)
	return int64(ans)
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}
	ans := 0
	var dfs func(u, p int) (mx, mx1 int)
	dfs = func(u, p int) (mx, mx1 int) {
		pm, pm1 := price[u], 0
		for _, v := range adj[u] {
			if v != p {
				sm, sm1 := dfs(v, u)
				ans = max(ans, sm+pm1, pm+sm1)
				pm, pm1 = max(pm, price[u]+sm), max(pm1, price[u]+sm1)
			}
		}
		return pm, pm1
	}
	dfs(0, -1)
	return int64(ans)
}
