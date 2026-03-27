package main

import "maps"

func tallestBillboard(rods []int) int {
	dp := map[int]int{0: 0}
	for _, v := range rods {
		old := maps.Clone(dp)
		for d, l := range old {
			dp[d+v] = max(dp[d+v], l+v)
			dp[d-v] = max(dp[d-v], l+v)
		}
	}
	return dp[0] / 2
}
