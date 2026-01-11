package greedy

func maxSubarraySumCircular(nums []int) int {
	mn, mx, sum, ans := 0, 0, 0, int(-2e9)
	for _, v := range nums {
		sum += v
	}
	for i, v := range nums {
		mn, mx = v+min(0, mn), v+max(0, mx)
		ans = max(ans, mx)
		if i < len(nums)-1 { // bug: remain at least one num
			ans = max(ans, sum-mn)
		}
	}
	return ans
}
