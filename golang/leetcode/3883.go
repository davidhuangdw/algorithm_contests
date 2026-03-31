package main

func countArrays(digitSum []int) int {
	M, mx := int(1e9+7), 5001

	ds := make([]int, mx)
	for i := 1; i < mx; i++ {
		ds[i] = ds[i/10] + i%10
	}

	dp := make([]int, mx)
	for v := 0; v < mx; v++ {
		if ds[v] == digitSum[0] {
			dp[v] = 1
		}
	}

	for _, tar := range digitSum[1:] {
		ndp, pre := make([]int, mx), 0
		for v := 0; v < mx; v++ {
			pre = (pre + dp[v]) % M
			if ds[v] == tar {
				ndp[v] = pre
			} else {
				ndp[v] = 0
			}
		}
		dp, ndp = ndp, dp
	}

	ans := 0
	for _, c := range dp {
		ans = (ans + c) % M
	}
	return ans
}
