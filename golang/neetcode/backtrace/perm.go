package backtrace

func permute(nums []int) [][]int {
	ans := make([][]int, 0)

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(0)

	return ans
}
