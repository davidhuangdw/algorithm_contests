package main

import (
	"fmt"
	"io"
)

func getCuts(n int, adj [][]int) []int {
	time, dfn, low := 0, make([]int, n+1), make([]int, n+1)
	isCuts := make(map[int]bool)
	var getLow func(u, par int) int
	getLow = func(u, par int) int {
		time++
		dfn[u] = time
		low[u] = dfn[u]
		child := 0 // bug: use childCnt instead of len(adj[root]) to check
		for _, v := range adj[u] {
			if v == par {
				continue
			}
			if dfn[v] == 0 {
				child++
				low[u] = min(low[u], getLow(v, u))
				if low[v] >= dfn[u] && par != -1 {
					isCuts[u] = true
				}
			} else {
				low[u] = min(low[u], dfn[v])
			}
		}
		if par == -1 && child > 1 {
			isCuts[u] = true
		}

		return low[u]
	}
	for u := 1; u <= n; u++ {
		if dfn[u] == 0 {
			getLow(u, -1)
		}
	}

	cuts := make([]int, 0)
	for u := 1; u <= n; u++ {
		if isCuts[u] {
			cuts = append(cuts, u)
		}
	}
	return cuts
}

func P3388_undirected_cut_vertex(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		if u != v {
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
	}

	cuts := getCuts(n, adj)
	fmt.Fprintln(OUT, len(cuts))
	for i, u := range cuts {
		fmt.Fprint(OUT, u)
		if i != len(cuts)-1 {
			fmt.Fprint(OUT, " ")
		}
	}
	fmt.Fprintln(OUT)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3388_undirected_cut_vertex(IN, OUT)
//}
