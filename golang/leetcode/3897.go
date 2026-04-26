package main

import "slices"

func maxValue(nums1 []int, nums0 []int) int {
	one, seg := 0, make([][2]int, 0, len(nums1))
	for i := range nums1 {
		if nums0[i] == 0 {
			one += nums1[i]
		} else {
			seg = append(seg, [2]int{nums0[i], nums1[i]})
		}
	}
	slices.SortFunc(seg, func(a, b [2]int) int { // both has zero
		if a[1] != b[1] {
			return b[1] - a[1]
		}
		return a[0] - b[0]
	})
	ans, M := 0, int(1e9+7)
	for range one {
		ans = (ans*2 + 1) % M
	}
	for _, s := range seg {
		for d := 1; d >= 0; d-- {
			for range s[d] {
				ans = (ans*2 + d) % M
			}
		}
	}
	return ans
}
