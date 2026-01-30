package main

import (
	"fmt"
	"io"
)

func buildDirectedSccGraph(n int, wei []int, adj [][]int) ([]int, []map[int]bool) {
	belong, sccSum := make([]int, n+1), make([]int, 0)
	time, dfn, low, que, qued := 0, make([]int, n+1), make([]int, n+1), make([]int, 0), make([]bool, n+1)
	var getLow func(u int) int
	getLow = func(u int) int { // lowest ancestor/not-poped-brother by at most one back-jump
		que = append(que, u)
		time++
		dfn[u], low[u], qued[u] = time, time, true
		for _, v := range adj[u] {
			if dfn[v] == 0 {
				low[u] = min(low[u], getLow(v))
			} else if qued[v] { // bug: only for ancestor/not-poped-brother(-> handle by LCA)
				low[u] = min(low[u], dfn[v])
			}
		}
		if low[u] >= dfn[u] {
			sccCnt, sum := len(sccSum), 0
			for {
				v := que[len(que)-1]
				que = que[:len(que)-1]
				qued[v] = false
				belong[v] = sccCnt
				sum += wei[v]
				if v == u {
					break
				}
			}
			sccSum = append(sccSum, sum)
		}
		return low[u]
	}
	for u := 1; u <= n; u++ {
		if dfn[u] == 0 {
			getLow(u)
		}
	}

	sccAdj := make([]map[int]bool, len(sccSum))
	for u := range sccAdj {
		sccAdj[u] = make(map[int]bool)
	}
	for u, es := range adj {
		for _, v := range es {
			sccAdj[belong[u]][belong[v]] = true
		}
	}
	return sccSum, sccAdj
}

func getMaxDirectedPath(n int, wei []int, adj [][]int) int {
	sccSum, sccAdj := buildDirectedSccGraph(n, wei, adj)
	dp, mx := make([]int, len(sccSum)), 0
	for u, s := range sccSum {
		for v := range sccAdj[u] {
			dp[u] = max(dp[u], dp[v])
		}
		dp[u] += s
		mx = max(mx, dp[u])
	}
	return mx
}

func P3387_directed_scc(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj, wei := make([][]int, n+1), make([]int, n+1)
	for u := 1; u <= n; u++ {
		fmt.Fscan(IN, &wei[u])
	}
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		if u != v {
			adj[u] = append(adj[u], v)
		}
	}

	fmt.Fprintln(OUT, getMaxDirectedPath(n, wei, adj))
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3387_directed_scc(IN, OUT)
//}
