package backtrace

import "slices"

func combinationSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)

	slices.Sort(nums)
	var cur []int
	fail := make(map[[2]int]bool)

	var dfs func(i, sum int) bool
	dfs = func(i, sum int) (succ bool) {
		defer func() {
			if !succ {
				fail[[2]int{i, sum}] = true // cache fail
			}
		}()

		if sum == target {
			ans = append(ans, append([]int{}, cur...))
			return true
		}
		if i >= len(nums) || sum > target ||
			sum+nums[i] > target || // early stop since (sorted)nums[] after i will also >target
			fail[[2]int{i, sum}] { // due to failed-tried before
			return false
		}

		//ignore i
		succ = dfs(i+1, sum) || succ

		// use i
		sum += nums[i]
		if sum <= target {
			cur = append(cur, nums[i])
			succ = dfs(i, sum) || succ
			cur = cur[:len(cur)-1]
		}

		return succ
	}

	dfs(0, 0)
	return ans
}
