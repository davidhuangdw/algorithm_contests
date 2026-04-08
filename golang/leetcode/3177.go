package main

import (
	"slices"
)

func maximumLength(nums []int, k int) int {
	n := len(nums)
	last, pos := make([]int, n), make(map[int]int)
	for i, v := range nums {
		if j, ok := pos[v]; ok {
			last[i] = j
		} else {
			last[i] = -1
		}
		pos[v] = i
	}

	dp := make([]int, n)
	for range k + 1 {
		var preMx int
		for i := range dp {
			res := preMx + 1
			if last[i] >= 0 {
				res = max(res, dp[last[i]]+1)
			}
			dp[i], preMx = res, max(preMx, dp[i])
		}
		// fmt.Println(dp)
	}
	return slices.Max(dp)
}
