package main

import (
	"fmt"
	"io"
)

func countTriangle(n int, edges [][2]int) (cnt int) {
	deg := make([]int, n+1)
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
	}
	adj := make([]map[int]bool, n+1)
	for i := range adj {
		adj[i] = make(map[int]bool)
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		if deg[u] > deg[v] || (deg[u] == deg[v] && u > v) { // degree-ordered edge direction
			u, v = v, u
		}
		adj[u][v] = true
	}

	for u := 1; u <= n; u++ {
		for v := range adj[u] {
			for w := range adj[v] {
				if adj[u][w] {
					cnt++
				}
			}
		}
	}
	return cnt
}

func P1989_count_triangle(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	edges := make([][2]int, m)
	for i := range edges {
		fmt.Fscan(IN, &edges[i][0], &edges[i][1])
	}

	ans := countTriangle(n, edges)
	fmt.Fprintln(OUT, ans)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P1989_count_triangle(IN, OUT)
//}
