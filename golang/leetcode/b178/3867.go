package main

import "slices"

func gcdSum(nums []int) int64 {
	gcd := func(a, b int) int {
		for a > 0 {
			a, b = b%a, a
		}
		return b
	}
	n := len(nums)
	a, mx := make([]int, n), -1
	for i, v := range nums {
		mx = max(mx, v)
		a[i] = gcd(mx, v)
	}
	slices.Sort(a)
	var ans int64
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		ans += int64(gcd(a[l], a[r]))
	}
	return ans
}
