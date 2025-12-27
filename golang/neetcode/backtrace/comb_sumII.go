package backtrace

import "slices"

func combinationSum2(nums []int, target int) (ans [][]int) {
	ans = make([][]int, 0)
	slices.Sort(nums)

	var cur []int
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, cur...))
			return
		}
		if i >= len(nums) || sum > target || sum+nums[i] > target {
			return
		}

		j := i + 1
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}

		dfs(j, sum)
		for k := i + 1; k <= j; k++ {
			sum += nums[i]
			cur = append(cur, nums[i])
			dfs(j, sum)
		}
		cur = cur[:len(cur)-(j-i)]
	}
	dfs(0, 0)
	return ans
}
