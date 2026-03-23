package main

import "slices"

func gcdSort(nums []int) bool {
	n := len(nums)
	fa := make(map[int]int)
	var find func(u int) int
	find = func(u int) int {
		if _, ok := fa[u]; !ok {
			fa[u] = u
		}
		if fa[u] != u {
			fa[u] = find(fa[u])
		}
		return fa[u]
	}
	merge := func(u, v int) {
		u, v = find(u), find(v)
		if u != v {
			fa[u] = v
		}
	}

	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	factor := make(map[int]int)
	for v, ind := range pos {
		lastP, x := -1, v
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				lastP = p
				merge(n+p, ind[0])
				for ; x%p == 0; x /= p {
				}
			}
		}
		if x > 1 {
			lastP = x
			merge(n+x, ind[0])
		}
		factor[v] = lastP
		for _, j := range ind[1:] {
			merge(j, ind[0])
		}
	}

	slices.Sort(nums)
	for i, v := range nums {
		if find(i) != find(n+factor[v]) {
			return false
		}
	}
	return true
}
