package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type RMQ struct {
	rg [][]int
}

func NewRMQ(nums []int) *RMQ {
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
	return &RMQ{rg: rg}
}

func (rm *RMQ) query(l, r int) int {
	d, k := r+1-l, 0
	for 1<<k <= d {
		k++
	}
	k--
	return max(rm.rg[l][k], rm.rg[r+1-(1<<k)][k])
}

func P3865_rmq(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &nums[i])
	}
	rmq := NewRMQ(nums)

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(IN, &l, &r)
		fmt.Fprintln(OUT, rmq.query(l-1, r-1))
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
