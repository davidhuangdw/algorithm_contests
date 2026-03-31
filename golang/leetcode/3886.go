package main

import "slices"

func sortableIntegers(nums []int) int {
	n := len(nums)
	ord := make([]int, n)
	copy(ord, nums)
	slices.Sort(ord)

	ans := 0
	for k := 1; k <= n; k++ {
		if n%k != 0 {
			continue
		}

		succ := true
		for i := 0; i < n && succ; i += k {
			down, pos := 0, -1
			for j := 0; j < k-1; j++ {
				if nums[i+j] > nums[i+j+1] {
					down++
					pos = j
				}
			}
			if down > 1 {
				succ = false
				break
			}

			st := 0
			if down == 1 {
				st = pos + 1
			}
			for j := 0; j < k; j++ {
				if nums[i+(st+j)%k] != ord[i+j] {
					succ = false
					break
				}
			}
		}
		if succ {
			ans += k
		}
	}
	return ans
}
