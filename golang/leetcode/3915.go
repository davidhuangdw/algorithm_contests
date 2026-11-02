package main

import "slices"

func maxAlternatingSum(nums []int, k int) int64 {
	n := len(nums)
	ord := make([]int, n)
	for i := range ord {
		ord[i] = i
	}
	slices.SortFunc(ord, func(a, b int) int {
		return nums[a] - nums[b]
	})
	rank, R := make([]int, n), 0
	for i, v := range ord {
		if !(i-1 >= 0 && nums[ord[i-1]] == nums[v]) {
			R++
		}
		rank[v] = R
	}

	update := func(bit []int, i, v int) {
		for ; i <= R; i += i & -i {
			bit[i] = max(bit[i], v)
		}
	}
	preMax := func(bit []int, i int) int {
		res := 0
		for ; i > 0; i &= i - 1 {
			res = max(res, bit[i])
		}
		return res
	}
	revRank := func(i int) int {
		return R + 1 - rank[i]
	}

	uBIT, dBIT := make([]int, R+1), make([]int, R+1)
	up, down, ans := make([]int, n), make([]int, n), 0
	for i := range n {
		if i-k >= 0 {
			update(uBIT, revRank(i-k), up[i-k])
			update(dBIT, rank[i-k], down[i-k])
		}
		up[i] = nums[i] + preMax(dBIT, rank[i]-1)
		down[i] = nums[i] + preMax(uBIT, revRank(i)-1)
		ans = max(ans, up[i], down[i])
	}
	return int64(ans)
}
