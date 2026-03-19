package main

import "slices"

func maxCapacity(costs []int, capacity []int, budget int) int {
	n := len(costs)
	ind := make([]int, n)
	for i := range ind {
		ind[i] = i
	}
	slices.SortFunc(ind, func(i, j int) int {
		return costs[i] - costs[j]
	})
	ans, mx := 0, 0
	for l, r := 0, n-1; l <= r; r-- {
		for l < r && costs[ind[l]]+costs[ind[r]] < budget {
			ans = max(ans, mx+capacity[ind[l]])
			mx = max(mx, capacity[ind[l]])
			l++
		}
		if costs[ind[r]] < budget {
			ans = max(ans, capacity[ind[r]]+mx)
		}
	}
	return ans
}
