package dp

import "slices"

func combinationSum4(nums []int, target int) int {
	slices.Sort(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		for _, v := range nums {
			if j-v < 0 {
				break
			}
			dp[j] += dp[j-v]
		}
	}
	return dp[target]
}
