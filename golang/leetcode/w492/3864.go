package main

// q4
func minCost(s string, encCost int, flatCost int) int64 {
	n := len(s)

	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + int(s[i]-'0')
	}

	mem := make(map[[2]int]int64)

	var dp func(int, int) int64
	dp = func(s, l int) int64 {
		if _, ok := mem[[2]int{s, l}]; !ok {
			x := pre[s+l] - pre[s]
			var ans int64
			if x == 0 {
				ans = int64(flatCost)
			} else {
				ans = int64(l) * int64(x) * int64(encCost)
			}

			if l > 1 && l%2 == 0 {
				half := l / 2
				ans = min(ans, dp(s, half)+dp(s+half, half))
			}

			mem[[2]int{s, l}] = ans
		}
		return mem[[2]int{s, l}]
	}
	return dp(0, n)
}
