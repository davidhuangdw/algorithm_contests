package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func dinicMaxFlow(s, t int, adj []map[int]int) int {
	n, inf, ans := len(adj), int(1e9), 0
	dis, adjV, now := make([]int, n), make([][]int, n), make([]int, n)
	bfsDis := func() bool {
		for u := range dis {
			dis[u], now[u], adjV[u] = inf, 0, nil
		}
		dis[s] = 0
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for v, w := range adj[u] {
				if w > 0 && dis[u]+1 < dis[v] {
					adjV[u] = append(adjV[u], v)
					dis[v], q = dis[u]+1, append(q, v)
					if v == t {
						return true
					}
				}
			}
		}
		return false
	}
	var dfsDrain func(u, in int) int
	dfsDrain = func(u, in int) int {
		if u == t {
			return in
		}
		used := 0
		for ; in > 0 && now[u] < len(adjV[u]); now[u]++ {
			v := adjV[u][now[u]]
			if !(dis[v] == dis[u]+1 && adj[u][v] > 0) {
				continue
			}
			out := dfsDrain(v, min(in, adj[u][v]))
			if out == 0 {
				continue
			}
			used += out
			in -= out
			adj[u][v] -= out
			adj[v][u] += out
			if in == 0 {
				break // 'u->v' is reusable: should keep now[u]
			}
		}
		if used == 0 {
			dis[u] = inf // cut branch
		}
		return used
	}

	for bfsDis() {
		ans += dfsDrain(s, inf)
	}
	return ans
}

func P3376_max_flow(IN io.Reader, OUT io.Writer) {
	var n, m, s, t int
	fmt.Fscan(IN, &n, &m, &s, &t)
	adj := make([]map[int]int, n+1)
	for i := range adj {
		adj[i] = make(map[int]int)
	}
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(IN, &u, &v, &w)
		adj[u][v] += w
	}
	ans := dinicMaxFlow(s, t, adj)
	fmt.Fprintln(OUT, ans)
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P3376_max_flow(IN, OUT)
}
