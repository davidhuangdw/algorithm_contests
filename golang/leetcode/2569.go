package main

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n, res := len(nums1), 0
	for _, v := range nums2 {
		res += v
	}
	sm, lzy := make([]int, n*4), make([]int, n*4)
	var build func(i, l, r int) int
	build = func(i, l, r int) int {
		if l == r {
			sm[i] = nums1[l]
		} else {
			md := (l + r) >> 1
			sm[i] = build(i*2, l, md) + build(i*2+1, md+1, r)
		}
		return sm[i]
	}
	build(1, 0, n-1)
	var flip func(i, l, r, ql, qr int)
	push := func(i, l, r int) {
		if lzy[i] == 1 {
			lc, rc, md := i*2, i*2+1, (l+r)>>1
			flip(lc, l, md, l, r)
			flip(rc, md+1, r, l, r)
			lzy[i] = 0
		}
	}
	flip = func(i, l, r, ql, qr int) {
		if r < ql || qr < l {
			return
		}
		if ql <= l && r <= qr {
			sm[i] = r + 1 - l - sm[i]
			lzy[i] ^= 1
			return
		}
		push(i, l, r)
		lc, rc, md := i*2, i*2+1, (l+r)>>1
		flip(lc, l, md, ql, qr)
		flip(rc, md+1, r, ql, qr)
		sm[i] = sm[lc] + sm[rc]
	}

	ans := make([]int64, 0)
	for _, q := range queries {
		switch q[0] {
		case 1:
			flip(1, 0, n-1, q[1], q[2])
		case 2:
			res += sm[1] * q[1]
		case 3:
			ans = append(ans, int64(res))
		}
	}
	return ans
}
