package backtrace

import "slices"

func permuteUnique1(nums []int) (ans [][]int) { // by cnt
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	var dfs func(pre []int)
	dfs = func(pre []int) {
		if len(pre) == len(nums) {
			ans = append(ans, append([]int{}, pre...))
			return
		}
		for v, c := range cnt {
			if c == 0 {
				continue
			}
			cnt[v]--
			dfs(append(pre, v))
			cnt[v]++
		}
	}
	dfs(nil)
	return ans
}

func permuteUnique(nums []int) (ans [][]int) { // by swap
	slices.Sort(nums)

	n := len(nums)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for j := i; j < n; j++ {
			if j > i && nums[j] == nums[i] { // choose nxt-larger
				continue
			}
			nums[i], nums[j] = nums[j], nums[i] // cur val swap nxt-larger val
			dfs(i + 1)
			// keep sorted by do nothing: since ascend swap nxt-larger val
		}
		// resort to sorted: now num[i]==largest(others are sorted)
		for j := n - 1; j > i; j-- {
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	dfs(0)
	return ans
}
