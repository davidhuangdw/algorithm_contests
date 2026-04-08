package main

func maximizeXor(nums []int, queries [][]int) []int {
	inf, B := 1<<60, 32
	type Node struct {
		ch  [2]*Node
		min int
	}
	New := func() *Node {
		return &Node{min: inf}
	}
	root := New()
	add := func(v int) {
		root.min = min(root.min, v)
		for b, p := B-1, root; b >= 0; b-- {
			d := v >> b & 1
			if p.ch[d] == nil {
				p.ch[d] = New()
			}
			p = p.ch[d]
			p.min = min(p.min, v)
		}
	}
	for _, v := range nums {
		add(v)
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		qx, qm := q[0], q[1]
		if root.min > qm {
			ans[i] = -1
			continue
		}
		p, res := root, 0
		for b := B - 1; b >= 0; b-- {
			dx := qx >> b & 1
			one := 1 ^ dx
			if p.ch[one] != nil && p.ch[one].min <= qm {
				res |= 1 << b
				p = p.ch[one]
			} else {
				p = p.ch[dx]
			}
		}
		ans[i] = res
	}
	return ans
}
