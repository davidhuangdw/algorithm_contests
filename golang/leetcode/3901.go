package main

func countGoodSubseq(nums []int, p int, queries [][]int) int {
	n := len(nums)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	// seg tree
	type node struct{ g, c int }
	t := make([]node, 4*n)
	var upd func(i, l, r, idx, v int)
	upd = func(i, l, r, idx, v int) {
		if l == r {
			if v%p == 0 {
				t[i] = node{v / p, 1}
			} else {
				t[i] = node{0, 0}
			}
			return
		}
		lc, rc, md := i*2, i*2+1, (l+r)/2
		if idx <= md {
			upd(lc, l, md, idx, v)
		} else {
			upd(rc, md+1, r, idx, v)
		}
		t[i] = node{gcd(t[lc].g, t[rc].g), t[lc].c + t[rc].c}
	}

	for i, v := range nums {
		upd(1, 0, n-1, i, v)
	}

	ans := 0
	for _, q := range queries {
		nums[q[0]] = q[1]
		upd(1, 0, n-1, q[0], q[1])

		if t[1].c > 0 && t[1].g == 1 {
			if t[1].c < n || n >= 7 {
				ans++
				continue
			}
			for i := range n {
				g := 0
				for j, v := range nums {
					if j != i {
						g = gcd(g, v/p)
					}
				}
				if g == 1 {
					ans++
					break
				}
			}
		}
	}
	return ans
}
