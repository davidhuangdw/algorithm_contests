package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func edgeBiconnected(adj [][]int) (bcc [][]int) {
	n := len(adj) - 1
	time, dfn, low := 1, make([]int, n+1), make([]int, n+1)
	var stk []int
	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		dfn[u], low[u], time = time, time, time+1
		stk = append(stk, u)
		for _, v := range adj[u] {
			if v == p {
				p = -1 // only ignore first back-edge, to allow multi-edge(e.g. 2-nodes-cycle)
				continue
			}
			if dfn[v] > 0 {
				low[u] = min(low[u], dfn[v])
			} else {
				tarjan(v, u)
				low[u] = min(low[u], low[v])
			}
		}
		if low[u] == dfn[u] {
			ui := len(stk) - 1
			for ; stk[ui] != u; ui-- {
			}
			bcc = append(bcc, append([]int{}, stk[ui:]...))
			stk = stk[:ui]
		}
	}
	for u := 1; u <= n; u++ {
		if dfn[u] == 0 {
			tarjan(u, -1)
		}
	}
	return bcc
}

func P8436_edge_biconnected(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	bcc := edgeBiconnected(adj)
	fmt.Fprintln(OUT, len(bcc))
	for _, c := range bcc {
		fmt.Fprint(OUT, len(c))
		for _, u := range c {
			fmt.Fprint(OUT, " ", u)
		}
		fmt.Fprintln(OUT)
	}

}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P8436_edge_biconnected(IN, OUT)
}
