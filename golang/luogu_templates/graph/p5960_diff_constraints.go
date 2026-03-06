package main

import (
	"fmt"
	"io"
)

func diffConstraints(n int, constraints [][3]int) []int {
	adj, x := make([][][2]int, n+1), make([]int, n+1)
	for _, c := range constraints {
		u, v, y := c[0], c[1], c[2]           // x[u] - x[v] <= y
		adj[v] = append(adj[v], [2]int{u, y}) // e(v, u) = y
	}

	// spfa
	qued, q := make([]bool, n+1), make([]int, 0, n)
	for u := 1; u <= n; u++ {
		q, qued[u], x[u] = append(q, u), true, 0
	}
	for rnd := 0; len(q) > 0; rnd++ {
		if rnd >= n {
			return nil
		}
		for k := len(q); k > 0; k-- {
			u := q[0]
			q, qued[u] = q[1:], false
			for _, e := range adj[u] {
				v, w := e[0], e[1]
				if x[u]+w < x[v] {
					x[v] = x[u] + w
					if !qued[v] {
						qued[v], q = true, append(q, v)
					}
				}
			}
		}
	}

	return x[1:]
}

func P5960_diff_constraints(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	constraints := make([][3]int, m)
	for i := range constraints {
		var u, v, y int // x[u] - x[v] <= y
		fmt.Fscan(IN, &u, &v, &y)
		constraints[i] = [3]int{u, v, y}
	}

	ans := diffConstraints(n, constraints)
	if ans == nil {
		fmt.Fprintln(OUT, "NO")
		return
	}
	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(OUT, " ")
		}
		fmt.Fprint(OUT, v)
	}
	fmt.Fprintln(OUT)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P5960_diff_constraints(IN, OUT)
//}
