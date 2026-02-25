package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type RMQByST struct {
	rg [][]int
}

func NewRMQByST(nums []int) *RMQByST {
	n := len(nums)
	rg := make([][]int, n)

	K := 0
	for 1<<K < n {
		K++
	}
	for i := 0; i < n; i++ {
		rg[i] = make([]int, K+1)
		rg[i][0] = nums[i]
	}
	for k := 0; k+1 <= K; k++ {
		for i := range rg {
			j := i + (1 << k)
			if j < n {
				rg[i][k+1] = max(rg[i][k], rg[j][k])
			} else {
				rg[i][k+1] = rg[i][k]
			}
		}
	}
	return &RMQByST{rg: rg}
}

func (rm *RMQByST) query(l, r int) int {
	d, k := r+1-l, 0
	for 1<<k <= d {
		k++
	}
	k--
	return max(rm.rg[l][k], rm.rg[r+1-(1<<k)][k])
}

func rmqByST(nums []int, queries [][2]int) []int {
	rmq := NewRMQByST(nums)
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = rmq.query(q[0], q[1])
	}
	return ans
}

func rmqByCartesianLCA(nums []int, queries [][2]int) []int { // O(n+m) by Cartesian-Tree + offline-tarjan-LCA
	n, m := len(nums), len(queries)
	ans := make([]int, m)

	// build cartesian tree
	chd := make([][2]int, n)
	for i := range chd {
		chd[i] = [2]int{-1, -1}
	}
	var q []int
	for i, v := range nums {
		pre := -1
		for len(q) > 0 && nums[q[len(q)-1]] < v {
			pre, q = q[len(q)-1], q[:len(q)-1]
		}
		chd[i][0] = pre
		if len(q) > 0 {
			chd[q[len(q)-1]][1] = i
		}
		q = append(q, i)
	}
	root := q[0]

	// offline tarjan LCA
	qs := make([][]int, n)
	for i, qu := range queries {
		a, b := qu[0], qu[1]
		qs[a] = append(qs[a], i)
		qs[b] = append(qs[b], i)
	}
	fa := make([]int, n) // union-find-set
	for i := range fa {
		fa[i] = i
	}
	var find func(u int) int
	find = func(u int) int {
		if fa[u] != u {
			fa[u] = find(fa[u])
		}
		return fa[u]
	}
	vis := make([]bool, n)
	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		if u < 0 {
			return
		}
		vis[u] = true
		for _, qi := range qs[u] {
			qu := queries[qi]
			v := qu[0] ^ qu[1] ^ u
			if vis[v] {
				ans[qi] = nums[find(v)]
			}
		}
		for _, v := range chd[u] {
			tarjan(v, u)
		}
		fa[u] = p
	}
	tarjan(root, -1)
	return ans
}

func P3865_rmq(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &nums[i])
	}

	queries := make([][2]int, m)
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(IN, &l, &r)
		queries[i][0] = l - 1
		queries[i][1] = r - 1
	}
	//ans := rmqByST(nums, queries)
	ans := rmqByCartesianLCA(nums, queries)

	for _, v := range ans {
		fmt.Fprintln(OUT, v)
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P3865_rmq(IN, OUT)
}
