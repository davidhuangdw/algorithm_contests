package main

import (
	"maps"
)

func goodSubtreeSum1(vals []int, par []int) int {
	B, ans := 1<<10, 0
	ch := make([][]int, len(vals))
	for i, p := range par {
		if p >= 0 {
			ch[p] = append(ch[p], i)
		}
	}

	var dfs func(u int) []int
	dfs = func(u int) []int {
		f, msk := make([]int, B), 0
		for v := vals[u]; v > 0; v /= 10 {
			d := v % 10
			if (1<<d)&msk > 0 {
				msk = 0
				break
			}
			msk |= 1 << d
		}
		if msk > 0 {
			f[msk] = vals[u]
		}

		for _, v := range ch[u] {
			for i, x := range dfs(v) {
				f[i] = max(f[i], x)
			}
		}

		for a := 1; a < B; a++ {
			for b := a & (a - 1); b > 0; b = (b - 1) & a {
				f[a] = max(f[a], f[b]+f[a^b])
			}
		}
		ans += f[B-1]
		return f
	}
	dfs(0)
	return ans % int(1e9+7)
}

func goodSubtreeSum(vals []int, par []int) int { // O(n*log* 2^D) by tree dsu
	ch, ans := make([][]int, len(vals)), 0
	for i, p := range par {
		if p >= 0 {
			ch[p] = append(ch[p], i)
		}
	}

	var dfs func(u int) (dp, vs map[int]int)
	dfs = func(u int) (f, vx map[int]int) {
		f, vx, msk := make(map[int]int), make(map[int]int), 0
		for v := vals[u]; v > 0; v /= 10 {
			d := v % 10
			if (1<<d)&msk > 0 {
				msk = 0
				break
			}
			msk |= 1 << d
		}
		if msk > 0 {
			f[msk] = vals[u]
			vx[msk] = vals[u]
		}

		for _, v := range ch[u] {
			fy, vy := dfs(v)
			// smaller vy put into larger vx
			if len(vy) > len(vx) {
				f, fy, vx, vy = fy, f, vy, vx
			}
			for msk, y := range vy {
				if y <= f[msk] {
					continue
				}
				vx[msk], f[msk] = y, y
				for mx, x := range maps.Clone(f) {
					if mx&msk == 0 {
						f[mx|msk] = max(f[mx|msk], x+y)
					}
				}
			}
		}
		mx := 0
		for _, x := range f {
			mx = max(mx, x)
		}
		ans += mx
		return f, vx
	}
	dfs(0)
	return ans % int(1e9+7)
}
