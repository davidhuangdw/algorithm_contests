package main

import (
	"fmt"
	"io"
)

func hasNegCycle(adj [][][]int) bool {
	n := len(adj)
	dis, qued := make([]int, n), make([]bool, n)
	for u := range dis {
		dis[u] = 2e9
	}

	dis[1], qued[1] = 0, true
	que, rnd := []int{1}, 0
	for ; len(que) > 0 && rnd < n; rnd++ {
		l := len(que)
		for i := 0; i < l; i++ {
			u := que[0]
			que, qued[u] = que[1:], false
			for _, e := range adj[u] {
				v, w := e[0], e[1]
				if dis[u]+w < dis[v] {
					dis[v] = dis[u] + w
					if !qued[v] {
						que, qued[v] = append(que, v), true
					}
				}
			}
		}
	}
	return rnd >= n
}

func P3385_neg_cycle(IN io.Reader, OUT io.Writer) {
	var T int
	fmt.Fscan(IN, &T)
	for ; T > 0; T-- {
		var n, m int
		fmt.Fscan(IN, &n, &m)
		adj := make([][][]int, n+1)
		for i := 0; i < m; i++ {
			var u, v, w int
			fmt.Fscan(IN, &u, &v, &w)
			adj[u] = append(adj[u], []int{v, w})
			if w >= 0 {
				adj[v] = append(adj[v], []int{u, w})
			}
		}
		if hasNegCycle(adj) {
			fmt.Fprintln(OUT, "YES")
		} else {
			fmt.Fprintln(OUT, "NO")
		}
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3385_neg_cycle(IN, OUT)
//}
