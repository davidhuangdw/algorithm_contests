package main

import (
	"sort"
)

func minStable(nums []int, maxC int) int {
	gcd := func(a, b int) int {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	n := len(nums)
	gs, dis := make([]int, n), make([]int, n)
	for i, x := range nums {
		gs[i], dis[i] = x, n-i
		if gs[i] == 1 {
			dis[i] = 0
		}
		for j := i - 1; j >= 0; j-- {
			g := gcd(x, gs[j])
			if gs[j] == g {
				break
			}
			gs[j] = g
			if gs[j] == 1 {
				dis[j] = i - j
			}
		}
	}
	return sort.Search(n+1, func(d int) bool {
		rem := maxC
		for i := 0; i < n; {
			if dis[i] <= d {
				i++
				continue
			}
			if rem <= 0 {
				return false
			}
			rem--
			i += d + 1
		}
		return true
	})
}
