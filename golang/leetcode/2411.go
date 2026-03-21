package main

import (
	"math/bits"
	"slices"
)

func smallestSubarrays1(nums []int) []int {
	B := max(bits.Len(uint(slices.Max(nums))), 1)
	pos, n := make([]int, B), len(nums)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- { // backward so that pos[] monotone-change
		v, j := nums[i], i
		for b := 0; b < B; b++ {
			if v>>b&1 > 0 {
				pos[b] = i
			}
			j = max(j, pos[b])
		}
		ans[i] = j + 1 - i
	}
	return ans
}

func smallestSubarrays(nums []int) []int { // logTrick
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = 1
		for j := i - 1; j >= 0 && (nums[j]|v) != nums[j]; j-- {
			nums[j] |= v
			ans[j] = i + 1 - j
		}
	}
	return ans
}
