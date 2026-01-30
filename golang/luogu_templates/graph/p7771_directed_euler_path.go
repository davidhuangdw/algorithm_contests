package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func getDirectedEulerPath(n int, adj [][]int) []int {
	diffDeg, m := make([]int, n+1), 0
	for u := range adj {
		slices.Sort(adj[u])
		m += len(adj[u])
		for _, v := range adj[u] {
			if u != v { // bug: could be same
				diffDeg[u]++
				diffDeg[v]--
			}
		}
	}

	st, out := 1, 0 // bug: maybe loop
	for u := 1; u <= n; u++ {
		if !(-1 <= diffDeg[u] && diffDeg[u] <= 1) {
			return nil
		}
		if diffDeg[u] == 1 {
			out++
			st = u
		}
	}

	if out > 1 {
		return nil
	}

	nxt, path := make([]int, n+1), make([]int, 0)
	var dfs func(u int)
	dfs = func(u int) {
		for nxt[u] < len(adj[u]) {
			v := adj[u][nxt[u]]
			nxt[u]++
			dfs(v)
		}
		path = append(path, u)
	}
	dfs(st)
	if len(path)-1 != m {
		return nil
	}

	slices.Reverse(path)
	return path
}

func P7771_eulerPath(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		adj[u] = append(adj[u], v)
	}

	path := getDirectedEulerPath(n, adj)
	if path == nil {
		fmt.Fprintln(OUT, "No")
		return
	}
	for i, u := range path {
		if i != 0 {
			fmt.Fprint(OUT, " ")
		}
		fmt.Fprint(OUT, u)
	}
	fmt.Fprintln(OUT)
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P7771_eulerPath(IN, OUT)
}
