package main

func minimumScore1(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}
	sub := make([]int, n)
	var dfs func(u, p int)
	dfs = func(u, p int) {
		sub[u] = nums[u]
		for _, v := range adj[u] {
			if v != p {
				dfs(v, u)
				sub[u] ^= sub[v]
			}
		}
	}
	dfs(0, -1)
	//fmt.Println(sub)
	var pre []int
	ans, path := int(1e10), make(map[int]bool)
	var calc func(u, p int)
	calc = func(u, p int) {
		for _, v := range pre {
			a, b, c := sub[u], sub[v], sub[0]^sub[v]
			if path[v] {
				b ^= a
			} else {
				c ^= a
			}
			//fmt.Println(u, v, a, b, c)
			ans = min(ans, max(a, b, c)-min(a, b, c))
		}

		path[u] = true
		if u != 0 {
			pre = append(pre, u)
		}
		for _, v := range adj[u] {
			if v != p {
				calc(v, u)
			}
		}
		path[u] = false
	}
	calc(0, -1)
	return ans
}

func minimumScore(nums []int, edges [][]int) int {
	n, t := len(nums), 0
	adj := make([][]int, n)
	for _, e := range edges {
		for i, u := range e {
			adj[u] = append(adj[u], e[i^1])
		}
	}

	in, out, sub := make([]int, n), make([]int, n), make([]int, n)
	var dfs func(u, p int)
	dfs = func(u, p int) {
		t++
		sub[u], in[u] = nums[u], t
		for _, v := range adj[u] {
			if v != p {
				dfs(v, u)
				sub[u] ^= sub[v]
			}
		}
		out[u] = t
	}
	dfs(0, -1)

	ans := int(1e10)
	for u := 1; u < n; u++ {
		for v := u + 1; v < n; v++ {
			p, q := u, v
			if in[p] > in[q] {
				p, q = q, p
			}
			a, b, c := sub[0]^sub[p], sub[p], sub[q]
			if in[q] <= out[p] {
				b ^= c
			} else {
				a ^= c
			}
			ans = min(ans, max(a, b, c)-min(a, b, c))
		}
	}
	return ans
}
