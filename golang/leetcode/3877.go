package main

import "maps"

func minRemovals(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 0

	for _, x := range nums {
		pre := maps.Clone(dp)
		for y, k := range pre {
			v := x ^ y
			dp[v] = max(pre[v], k+1)
		}
	}
	if k, ok := dp[target]; ok {
		return len(nums) - k
	}
	return -1
}
