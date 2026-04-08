package main

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > n/2 {
		return -1
	}
	if k == 0 {
		return 0
	}

	cost, inf := make([]int, n), int(1e18)
	for i := 0; i < n; i++ {
		l, r := (i-1+n)%n, (i+1)%n
		cost[i] = max(0, max(nums[l], nums[r])+1-nums[i])
	}

	solve := func(a []int, k int) int {
		m := len(a)
		dp := make([]int, m)

		for t := 1; t <= k; t++ {
			ndp := make([]int, m)
			for i, c := range a {
				r := inf
				if i-1 >= 0 {
					r = ndp[i-1]
				}
				if t == 1 {
					r = min(r, c)
				}
				if i-2 >= 0 {
					r = min(r, dp[i-2]+c)
				}
				ndp[i] = r
			}
			dp = ndp
		}
		return dp[m-1]
	}
	return min(solve(cost[1:], k), solve(cost[:n-1], k))
}
