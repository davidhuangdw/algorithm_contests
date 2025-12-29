package dp

func maxProduct(nums []int) int {
	mn, mx, ans := 1, 1, nums[0]
	for _, v := range nums {
		mn, mx = min(mn*v, mx*v, v), max(mn*v, mx*v, v)
		ans = max(ans, mx)
	}
	return ans
}
