package backtrace

func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	var cur []int
	cur = make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			clone := append([]int{}, cur...)
			ans = append(ans, clone)
			return
		}
		dfs(i + 1)

		cur = append(cur, nums[i])
		dfs(i + 1)
		cur = cur[:len(cur)-1]
	}
	dfs(0)
	return ans
}

func subsets(nums []int) (ans [][]int) { // iter
	ans = append(ans, []int{})
	for _, v := range nums {
		n := len(ans)
		for i := 0; i < n; i++ {
			addV := append([]int(nil), ans[i]...)
			addV = append(addV, v)
			ans = append(ans, addV)
		}
	}
	return ans
}
