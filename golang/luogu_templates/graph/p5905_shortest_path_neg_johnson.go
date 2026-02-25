package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

func shortestPathJohnson(adj [][][2]int) [][]int {
	n := len(adj)
	// spfa
	h, rnd := make([]int, n), 0
	que, qued := make([]int, n), make([]bool, n)
	for i := range que {
		que[i], h[i], qued[i] = i, 0, true
	}
	for ; len(que) > 0; rnd++ {
		if rnd >= n {
			return nil // neg cycle
		}
		for k := len(que); k > 0; k-- {
			u := que[0]
			qued[u], que = false, que[1:]
			for _, e := range adj[u] {
				v, w := e[0], e[1]
				if h[u]+w < h[v] {
					h[v] = h[u] + w
					if !qued[v] {
						qued[v], que = true, append(que, v)
					}
				}
			}
		}
	}

	// dijkstra
	dis := make([][]int, n)
	for i := range dis {
		d := make([]int, n)
		dis[i] = d
		for j := range d {
			d[j] = -1
		}
		d[i] = 0
		hq := NewHeap(func(a, b [2]int) bool {
			return a[0] < b[0]
		})
		heap.Push(hq, [2]int{0, i})
		for hq.Len() > 0 {
			pr := heap.Pop(hq).([2]int)
			du, u := pr[0], pr[1]
			for _, e := range adj[u] {
				v, w := e[0], e[1]
				w += h[u] - h[v] // become positive
				if d[v] < 0 || du+w < d[v] {
					d[v] = du + w
					heap.Push(hq, [2]int{d[v], v})
				}
			}
		}
	}

	// restore weight
	for u := range dis {
		for v, d := range dis[u] {
			if d >= 0 {
				dis[u][v] = d + h[v] - h[u]
			} else {
				dis[u][v] = 1e9 // unreachable
			}
		}
	}
	return dis
}

func P5905_shortest_path_neg_johnson(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	adj := make([][][2]int, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(IN, &u, &v, &w)
		adj[u-1] = append(adj[u-1], [2]int{v - 1, w})
	}

	dis := shortestPathJohnson(adj)
	if dis == nil {
		fmt.Fprintln(OUT, "-1")
		return
	}
	for _, row := range dis {
		ans := 0
		for j, d := range row {
			ans += (j + 1) * d
		}
		fmt.Fprintln(OUT, ans)
	}

}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P5905_shortest_path_neg_johnson(IN, OUT)
}
