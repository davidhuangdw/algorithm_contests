package main

import (
	"fmt"
	"io"
)

func bipartMaxMatch(n, m int, adj [][]int) int {
	ans := 0
	vis := make([]bool, n)
	match := make([]int, m)
	for i := range match {
		match[i] = -1
	}

	var extend func(l int) bool
	extend = func(l int) bool {
		if vis[l] {
			return false
		}
		vis[l] = true
		for _, r := range adj[l] {
			if match[r] == -1 || extend(match[r]) {
				match[r] = l
				return true
			}
		}
		return false
	}
	for l := 0; l < n; l++ {
		for i := range vis {
			vis[i] = false
		}
		if extend(l) {
			ans++
		}
	}
	return ans
}

func P3386_bipart_max_match(IN io.Reader, OUT io.Writer) {
	var n, m, e int
	fmt.Fscan(IN, &n, &m, &e)
	adj := make([][]int, n)
	exist := make(map[[2]int]bool)
	for i := 0; i < e; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		if !exist[[2]int{u, v}] {
			adj[u-1] = append(adj[u-1], v-1)
			exist[[2]int{u, v}] = true
		}
	}
	fmt.Fprintln(OUT, bipartMaxMatch(n, m, adj))
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3386_bipart_max_match(IN, OUT)
//}
