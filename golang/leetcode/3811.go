package main

func alternatingXOR(nums []int, target1 int, target2 int) int {
	n, M := len(nums), int(1e9+7)
	dp := [2]map[int]int{}
	for i := range dp {
		dp[i] = make(map[int]int)
		dp[i][0] = 1
	}
	tar := []int{target1, target2}
	cur := 0
	for i := n - 1; i >= 0; i-- {
		cur ^= nums[i]
		st1 := dp[1][cur^tar[0]] // bug: shouldn't immediately add to dp[0], will affect same-loop-read
		st2 := dp[0][cur^tar[1]]
		if i == 0 {
			return st1
		}
		dp[0][cur] = (dp[0][cur] + st1) % M
		dp[1][cur] = (dp[1][cur] + st2) % M
	}
	return -1
}
