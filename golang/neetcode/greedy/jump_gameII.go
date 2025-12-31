package greedy

func jump(nums []int) int {
	ans, l, r := 0, 0, 0
	for l <= r && r < len(nums)-1 {
		ans++
		old := r
		for i := l; i <= old; i++ {
			r = max(r, i+nums[i])
		}
		l = old + 1
	}
	return ans
}
