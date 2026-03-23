package main

import "slices"

func makeArrayIncreasing(a []int, b []int) int {
	inf, n, m := int(1e10), len(a), len(b)
	slices.Sort(b)

	type State struct {
		i, j, nxt int
	}
	dp := make(map[State]int)
	var dfs func(i, j, nxt int) int
	dfs = func(i, j, nxt int) int {
		if i < 0 {
			return 0
		}
		s := State{i, j, nxt}
		if _, ok := dp[s]; !ok {
			res := inf
			// keep a[i]
			if a[i] < nxt {
				res = min(res, dfs(i-1, j, a[i]))
			}
			if j >= 0 {
				// use b[j]
				if b[j] < nxt {
					res = min(res, dfs(i-1, j-1, b[j])+1)
				} else {
					res = min(res, dfs(i, j-1, nxt))
				}
			}
			dp[s] = res
		}
		return dp[s]
	}
	ans := dfs(n-1, m-1, inf)
	if ans == inf {
		return -1
	}
	return ans
}
