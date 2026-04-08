package main

import (
	"slices"
)

func minAbsDifference(nums []int, goal int) int {
	gen := func(a []int) []int {
		res := make([]int, 0, 1<<len(a))
		res = append(res, 0)
		for _, v := range a {
			t := res
			for _, x := range t {
				res = append(res, x+v)
			}
		}
		slices.Sort(res)
		return res
	}
	n := len(nums)
	a, b := gen(nums[:n/2]), gen(nums[n/2:])
	ans, r := int(1e18), len(b)-1
	for _, v := range a {
		for ; r >= 0 && v+b[r] > goal; r-- {
		}
		if r >= 0 {
			ans = min(ans, goal-(v+b[r]))
		}
		if r+1 < len(b) {
			ans = min(ans, v+b[r+1]-goal)
		}
	}
	return ans
}
