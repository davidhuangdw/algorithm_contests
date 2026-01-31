package main

import (
	"fmt"
	"io"
)

func queryLCA(n, root int, adj [][]int, queries [][2]int) []int {
	m := len(queries)
	ans := make([]int, m)
	qIdx := make([][]int, n+1)
	for j, q := range queries {
		for _, u := range q {
			qIdx[u] = append(qIdx[u], j)
		}
	}

	// union-find-set
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var getRoot func(u int) int
	getRoot = func(u int) int {
		if fa[u] != u {
			fa[u] = getRoot(fa[u])
		}
		return fa[u]
	}

	// tarjan dfs
	vis := make([]bool, n+1)
	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		vis[u] = true
		for _, qi := range qIdx[u] {
			q := queries[qi]
			v := q[0] ^ q[1] ^ u
			if vis[v] {
				ans[qi] = getRoot(v)
			}
		}
		for _, v := range adj[u] {
			if v != p {
				tarjan(v, u)
			}
		}
		fa[u] = p
	}
	tarjan(root, -1)
	return ans
}

func P3379_lca(IN io.Reader, OUT io.Writer) {
	var n, m, s, u, v int
	fmt.Fscan(IN, &n, &m, &s)
	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(IN, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	queries := make([][2]int, m)
	for i := range queries {
		fmt.Fscan(IN, &queries[i][0], &queries[i][1])
	}

	ans := queryLCA(n, s, adj, queries)
	for _, res := range ans {
		fmt.Fprintln(OUT, res)
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3379_lca(IN, OUT)
//}
