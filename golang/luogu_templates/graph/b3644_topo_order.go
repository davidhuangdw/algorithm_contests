package main

import (
	"fmt"
	"io"
	"slices"
)

func topoOrder(n int, adj [][]int) []int {
	ans, vis := make([]int, 0), make([]bool, n+1)
	var dfs func(u int)
	dfs = func(u int) {
		if vis[u] {
			return
		}
		vis[u] = true
		for _, v := range adj[u] {
			dfs(v)
		}
		ans = append(ans, u)
	}
	for u := 1; u <= n; u++ {
		dfs(u)
	}
	slices.Reverse(ans)
	return ans
}

func P3644_topo_order(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for {
			var j int
			fmt.Fscan(IN, &j)
			if j == 0 {
				break
			}
			adj[i] = append(adj[i], j)
		}
	}
	for i, u := range topoOrder(n, adj) {
		if i > 0 {
			fmt.Fprint(OUT, " ")
		}
		fmt.Fprint(OUT, u)
	}
	fmt.Fprintln(OUT)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3644_topo_order(IN, OUT)
//}
