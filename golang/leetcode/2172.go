package main

import "slices"

func maximumANDSum(nums []int, numSlots int) int {
	msk := 1
	for i := 0; i < numSlots; i++ {
		msk *= 3
	}

	dp := make([]int, msk)
	for st := 0; st < msk-1; st++ {
		cnt, k := make([]int, numSlots), 0
		for i, x := 0, st; x > 0; i, x = i+1, x/3 {
			cnt[i] = x % 3
			k += cnt[i]
		}
		if k >= len(nums) {
			continue
		}

		for i, b := 1, 1; i <= numSlots; i, b = i+1, b*3 {
			if cnt[i-1] < 2 {
				dp[st+b] = max(dp[st+b], dp[st]+i&nums[k])
			}
		}
	}
	return slices.Max(dp)
}
