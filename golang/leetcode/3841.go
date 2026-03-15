package main

import (
	"fmt"
	"strconv"
	"strings"
)

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	qLca := make([][]int, n) // u -> qi
	lca := make(map[int]int)
	qs := make([][3]int, len(queries))
	for qi, q := range queries {
		p := strings.Split(q, " ")
		if p[0][0] == 'u' {
			u, _ := strconv.Atoi(p[1])
			qs[qi] = [3]int{1, u, 1 << (p[2][0] - 'a')}
		} else {
			u, _ := strconv.Atoi(p[1])
			v, _ := strconv.Atoi(p[2])
			qs[qi] = [3]int{0, u, v}
			qLca[u] = append(qLca[u], qi)
			qLca[v] = append(qLca[v], qi)
		}
	}

	// fenwick tree
	BIT := make([]int, n+1)
	update := func(t, v int) {
		for ; 0 < t && t <= n; t += t & -t {
			BIT[t] ^= v
		}
	}
	preSum := func(t int) (res int) {
		for ; t > 0; t &= t - 1 {
			res ^= BIT[t]
		}
		return res
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(u int) int
	find = func(u int) int {
		if fa[u] != u {
			fa[u] = find(fa[u])
		}
		return fa[u]
	}

	time, st, ed := 0, make([]int, n), make([]int, n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		time++
		st[u] = time
		// tarjan lca
		for _, qi := range qLca[u] {
			v := u ^ qs[qi][1] ^ qs[qi][2]
			if st[v] > 0 {
				lca[qi] = find(v)
			}
		}

		// recur
		for _, v := range adj[u] {
			if v != p {
				dfs(v, u)
			}
		}
		fa[u] = p // merge
		ed[u] = time
	}
	dfs(0, -1)

	cur := make([]int, n)
	for u, ch := range []byte(s) {
		cur[u] = 1 << (ch - 'a')
		update(st[u], cur[u])
		update(ed[u]+1, cur[u]) // bug: forget restore outside-sub-tree
	}

	var ans []bool
	for qi, q := range qs {
		if q[0] == 0 {
			u, v := q[1], q[2]
			sum := preSum(st[u]) ^ preSum(st[v]) ^ cur[lca[qi]]
			ans = append(ans, sum&(sum-1) == 0)
		} else {
			u, val := q[1], q[2]
			update(st[u], cur[u]^val) // update t == affect [t, inf]
			update(ed[u]+1, cur[u]^val)
			cur[u] = val
		}
	}
	return ans
}

func main() {
	result := palindromePath(4, [][]int{{0, 1}, {0, 2}, {0, 3}}, "abca", []string{"query 1 2", "update 0 b", "query 2 3", "update 3 a", "query 1 3"})
	fmt.Println(result)
}
