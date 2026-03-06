package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
)

func moQueries(nums []int, q [][2]int) []int {
	mx, n, m := 0, len(nums), len(q)
	S := int(math.Sqrt(float64(n)))
	for _, v := range nums {
		mx = max(mx, v)
	}
	ord := make([]int, m)
	for i := range ord {
		ord[i] = i
	}
	slices.SortFunc(ord, func(i, j int) int {
		bi, bj := q[i][0]/S, q[j][0]/S
		if bi == bj {
			return q[i][1] - q[j][1]
		}
		return bi - bj
	})

	ans := make([]int, m)
	cur, sum, l, r := make([]int, mx+1), 0, 1, 0
	add := func(i, sign int) {
		v := nums[i]
		sum += cur[v]*2*sign + 1 // -cur[v]^2 +(cur[v]+sign)^2 == +(2*cur[v]*sign+1)
		cur[v] += sign
	}
	for _, qi := range ord {
		ql, qr := q[qi][0], q[qi][1]
		for l != ql {
			if l < ql {
				add(l, -1)
				l++
			} else {
				l--
				add(l, 1)
			}
		}
		for r != qr {
			if r < qr {
				r++
				add(r, 1)
			} else {
				add(r, -1)
				r--
			}
		}
		ans[qi] = sum
	}
	return ans
}

func P2709_mo_queries(IN io.Reader, OUT io.Writer) {
	var n, m, k int
	fmt.Fscan(IN, &n, &m, &k)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(IN, &nums[i])
	}
	queries := make([][2]int, m)
	for i := range queries {
		fmt.Fscan(IN, &queries[i][0], &queries[i][1])
	}
	ans := moQueries(nums, queries)
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
	P2709_mo_queries(IN, OUT)
}
