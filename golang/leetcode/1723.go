package main

func minimumTimeRequired(jobs []int, k int) int {
	U := 1 << len(jobs)
	t := make([]int, U)
	for i, v := range jobs {
		t[1<<i] = v
	}
	dp := make([][]int, k)
	dp[0] = make([]int, U)
	for b := 1; b < U; b++ {
		t[b] = t[b&(b-1)] + t[b&-b]
		dp[0][b] = t[b]
	}
	for i := 1; i < k; i++ {
		dp[i] = make([]int, U)
		for b := range dp[i] {
			mn := dp[i-1][b]
			for x := b; x > b-x; x = (x - 1) & b {
				mn = min(mn, max(t[x], dp[i-1][b-x]))
			}
			dp[i][b] = mn
		}
	}
	return dp[k-1][U-1]
}
