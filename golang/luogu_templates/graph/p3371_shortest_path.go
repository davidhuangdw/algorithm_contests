package main

import (
	"container/heap"
	"fmt"
	"io"
)

func spfa(n, s int, adj [][][2]int) []int {
	dis := make([]int, n+1)
	for i := range dis {
		dis[i] = 1<<31 - 1
	}
	dis[0], dis[s] = 0, 0
	qued, toExtend := make([]bool, n+1), []int{s}
	for len(toExtend) > 0 {
		u := toExtend[0]
		toExtend, qued[u] = toExtend[1:], false
		for _, e := range adj[u] {
			v, w := e[0], e[1]
			if dis[u]+w < dis[v] {
				dis[v] = dis[u] + w
				if !qued[v] {
					toExtend = append(toExtend, v)
					qued[v] = true
				}
			}
		}
	}
	return dis
}

func dijkstra(n, s int, adj [][][2]int) []int {
	dis, done := make([]int, n+1), make([]bool, n+1)
	for i := range dis {
		dis[i] = 1<<31 - 1
	}
	dis[0], dis[s] = 0, 0
	hp := NewHeap(func(a, b [2]int) bool {
		return a[0] < b[0]
	})
	heap.Push(hp, [2]int{0, s})
	for hp.Len() > 0 {
		p := heap.Pop(hp).([2]int)
		d, u := p[0], p[1]
		if done[u] {
			continue
		}
		dis[u] = d
		done[u] = true
		for _, e := range adj[u] {
			v, w := e[0], e[1]
			if !done[v] && d+w < dis[v] {
				dis[v] = d + w
				heap.Push(hp, [2]int{d + w, v})
			}
		}
	}
	return dis
}

func P3371_shortest_path(IN io.Reader, OUT io.Writer) {
	var n, m, s int
	fmt.Fscan(IN, &n, &m, &s)
	adj := make([][][2]int, n+1)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(IN, &u, &v, &w)
		adj[u] = append(adj[u], [2]int{v, w})
	}

	//ans := spfa(n, s, adj)[1:]
	ans := dijkstra(n, s, adj)[1:]
	for i, d := range ans {
		if i > 0 {
			fmt.Fprint(OUT, " ")
		}
		fmt.Fprint(OUT, d)
	}
	fmt.Fprintln(OUT)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3371_shortest_path(IN, OUT)
//}
