package backtrace

import "slices"

func subsetsWithDup1(nums []int) [][]int { // dfs for each distinct-val: choose count 0~k
	ans := make([][]int, 0)
	slices.Sort(nums)
	var cur []int
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			ans = append(ans, append([]int{}, cur...))
			return
		}
		j := i + 1
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		dfs(j)
		for k := 0; k < j-i; k++ {
			cur = append(cur, nums[i])
			dfs(j)
		}
		cur = cur[:len(cur)-(j-i)]
	}
	dfs(0)
	return ans
}

func subsetsWithDup2(nums []int) [][]int { // dfs for each pos: choose a distinct val
	ans := make([][]int, 0)
	slices.Sort(nums)
	var dfs func(int, []int)
	dfs = func(i int, cur []int) {
		ans = append(ans, append([]int{}, cur...))
		for j := i; j < len(nums); j++ { // choose last distinct val for pos 'len(cur)+1'
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			dfs(j+1, append(cur, nums[j]))
		}
	}

	dfs(0, nil)
	return ans
}

func subsetsWithDup(nums []int) [][]int { // iter with O(1) extra space
	ans := [][]int{[]int{}}
	slices.Sort(nums)

	for i := 0; i < len(nums); {
		j := i + 1
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		m := len(ans)
		for t := 0; t < m; t++ {
			cur := append([]int{}, ans[t]...)
			for u := 0; u < j-i; u++ {
				cur = append(cur, nums[i]) // ok to reuse 'cur' for same values, since diff value will clone
				ans = append(ans, cur)
			}
		}
		i = j
	}

	return ans
}
