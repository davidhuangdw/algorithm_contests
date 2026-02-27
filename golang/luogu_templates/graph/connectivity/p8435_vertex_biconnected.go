package main

import (
	"fmt"
	"io"
)

func vertexBiconnected(adj [][]int) [][]int {
	n := len(adj) - 1

	time, dfn, low := 1, make([]int, n+1), make([]int, n+1)
	var stk []int
	var bcc [][]int // every vertex belong to a (lowest)cut's bcc(bi-connected-component)
	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		dfn[u], low[u], stk = time, time, append(stk, u)
		time++
		for _, v := range adj[u] { // allow parent
			if v == p {
				continue
			}
			if dfn[v] > 0 {
				low[u] = min(low[u], dfn[v])
				continue
			}
			tarjan(v, u)
			low[u] = min(low[u], low[v])
			if low[v] >= dfn[u] { // if no lower, then must be poped (e.g. just single edge)
				// u is cut (if not root)
				sub, j := []int{u}, 0
				for j = len(stk) - 1; stk[j] != v; j-- {
				}
				sub, stk = append(sub, stk[j:]...), stk[:j] // pop till v
				bcc = append(bcc, sub)
			}
		}
	}
	for u := 1; u <= n; u++ {
		if dfn[u] == 0 {
			tarjan(u, -1)
			if dfn[u] == time-1 { // single vertex graph
				bcc = append(bcc, []int{u})
			}
		}
	}

	return bcc
}

func P8435_vertex_biconnected(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	bcc := vertexBiconnected(adj)
	fmt.Fprintln(OUT, len(bcc))
	for _, sub := range bcc {
		fmt.Fprint(OUT, len(sub))
		for _, u := range sub {
			fmt.Fprint(OUT, " ", u)
		}
		fmt.Fprintln(OUT)
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P8435_vertex_biconnected(IN, OUT)
//}
