package main

func countSubstrings(s string) int64 {
	ans := 0
	var dp [10][10]int
	for _, c := range []byte(s) {
		var ndp [10][10]int
		di := int(c - '0')
		for d := 1; d <= 9; d++ {
			for r := 0; r < d; r++ {
				ndp[d][(r*10+di)%d] += dp[d][r]
			}
			ndp[d][di%d]++
		}
		if di > 0 {
			ans += ndp[di][0]
		}
		dp = ndp
	}
	return int64(ans)
}
