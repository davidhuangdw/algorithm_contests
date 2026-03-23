package main

import (
	"slices"
)

func makeParityAlternating(nums []int) []int {
	if len(nums) == 1 {
		return []int{0, 0}
	}
	gmn, gmx := slices.Min(nums), slices.Max(nums)
	check := func(d int) []int {
		var mn, mx int
		mn, mx = 2e9, -2e9
		cnt := 0
		for i, v := range nums {
			if (i+d)&1 != v&1 {
				cnt++
				if v == gmn {
					v++
				} else if v == gmx {
					v--
				}
			}
			mn = min(mn, v)
			mx = max(mx, v)
		}
		return []int{cnt, max(mx-mn, 1)}
	}
	a, b := check(0), check(1)
	if a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]) {
		return a
	}
	return b
}
