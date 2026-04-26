package main

import "sort"

func kthRemainingInteger(nums []int, queries [][]int) []int {
	pre := make([]int, len(nums)+1)
	for i, x := range nums {
		pre[i+1] = pre[i]
		if x%2 == 0 {
			pre[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r, k := q[0], q[1], q[2]
		ans[i] = 2 * sort.Search(k+r-l+2, func(md int) bool {
			if md == 0 {
				return false
			}
			i := sort.Search(r-l+1, func(j int) bool {
				return nums[l+j] > 2*md
			})
			return md-(pre[l+i]-pre[l]) >= k
		})
	}

	return ans
}
