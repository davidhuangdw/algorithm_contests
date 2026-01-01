package dp

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target < 0 {
		target = -target
	}
	if !(sum >= target && (sum-target)%2 == 0) {
		return 0
	}
	sum = (sum - target) / 2

	dp := make([]int, sum+1)
	dp[0] = 1
	for _, v := range nums {
		for t := sum; t-v >= 0; t-- {
			dp[t] += dp[t-v]
		}
	}
	return dp[sum]
}
