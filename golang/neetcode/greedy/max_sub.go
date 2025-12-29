package greedy

func maxSubArray(nums []int) int {
	INF := int(1e9)
	mx, cur := -INF, -INF
	for _, v := range nums {
		cur = v + max(0, cur)
		mx = max(mx, cur)
	}
	return mx
}
