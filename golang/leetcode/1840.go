package main

import (
	"slices"
)

func maxBuilding(n int, restrictions [][]int) int {
	rs := append(restrictions, []int{1, 0})
	slices.SortFunc(rs, func(a, b []int) int {
		return a[0] - b[0]
	})
	if len(rs) > 0 && rs[len(rs)-1][0] != n {
		rs = append(rs, []int{n, n - 1})
	}

	for i := 0; i+1 < len(rs); i++ {
		a, b := rs[i], rs[i+1]
		b[1] = min(b[1], a[1]+b[0]-a[0])
	}

	for i := len(rs) - 1; i-1 >= 0; i-- {
		a, b := rs[i-1], rs[i]
		a[1] = min(a[1], b[1]+b[0]-a[0])
	}

	//fmt.Println(rs)
	ans := 0
	for i := 0; i+1 < len(rs); i++ {
		a, b := rs[i], rs[i+1]
		ans = max(ans, (a[1]+b[1]+b[0]-a[0])/2)
	}
	return ans
}
