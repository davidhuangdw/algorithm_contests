package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type SegTree struct {
	n           int
	_sum, _lazy []int
}

func NewSegTree(n int, initNums []int) *SegTree { // [0, n)
	tr := &SegTree{
		n:     n,
		_sum:  make([]int, n*4),
		_lazy: make([]int, n*4),
	}
	var build func(i, il, ir int)
	build = func(i, il, ir int) {
		if il >= ir {
			if il == ir && il < len(initNums) {
				tr._sum[i] = initNums[il]
			}
			return
		}
		lc, rc, md := i<<1, i<<1+1, (il+ir)>>1
		build(lc, il, md)
		build(rc, md+1, ir)
		tr._sum[i] = tr._sum[lc] + tr._sum[rc]
	}
	build(1, 0, tr.n-1)
	return tr
}

func (se *SegTree) Add(l, r, v int) {
	se.add(1, 0, se.n-1, l, r, v)
}
func (se *SegTree) Sum(l, r int) int {
	return se.sum(1, 0, se.n-1, l, r)
}

func (se *SegTree) add(i, il, ir, l, r, v int) {
	if ir < l || r < il {
		return
	}
	if l <= il && ir <= r {
		se.fullAdd(i, il, ir, v)
		return
	}
	se.push(i, il, ir) // bug: forgot Push
	lc, rc, md := i<<1, i<<1+1, (il+ir)>>1
	se.add(lc, il, md, l, r, v)
	se.add(rc, md+1, ir, l, r, v)
	se._sum[i] = se._sum[lc] + se._sum[rc] // bug: forgot Update
}

func (se *SegTree) sum(i, il, ir, l, r int) int {
	if ir < l || r < il {
		return 0
	}
	if l <= il && ir <= r {
		return se._sum[i]
	}
	se.push(i, il, ir)
	lc, rc, md := i<<1, i<<1+1, (il+ir)>>1
	return se.sum(lc, il, md, l, r) + se.sum(rc, md+1, ir, l, r)
}

func (se *SegTree) push(i, il, ir int) {
	if se._lazy[i] == 0 || il == ir {
		return
	}
	lc, rc, md := i<<1, i<<1+1, (il+ir)>>1
	se.fullAdd(lc, il, md, se._lazy[i])
	se.fullAdd(rc, md+1, ir, se._lazy[i])
	se._lazy[i] = 0
}

func (se *SegTree) fullAdd(i, il, ir, add int) {
	sz := ir - il + 1
	se._sum[i] += add * sz
	se._lazy[i] += add
}

func P(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(IN, &nums[i])
	}

	fmt.Fscanln(IN)
	tr := NewSegTree(n+1, nums)
	for i := 0; i < m; i++ {
		var op, x, y, k int
		fmt.Fscanln(IN, &op, &x, &y, &k)
		switch op {
		case 1:
			tr.Add(x, y, k)
		case 2:
			fmt.Fprintln(OUT, tr.Sum(x, y))
		}
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P(IN, OUT)
}
