package main

func countSequences(nums []int, k int64) int {
	mem, n := make(map[[3]int]int), len(nums)
	var dp func(i, p, q int) int
	dp = func(i, p, q int) int {
		key := [3]int{i, p, q}
		if _, ok := mem[key]; !ok {
			var res int
			if i >= n {
				if p%q == 0 && p/q == int(k) {
					res = 1
				}
			} else {
				res = dp(i+1, p, q) + dp(i+1, p*nums[i], q) + dp(i+1, p, q*nums[i])
			}
			mem[key] = res
		}
		return mem[key]
	}
	return dp(0, 1, 1)
}
