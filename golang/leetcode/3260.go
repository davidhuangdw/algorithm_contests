package main

func largestPalindrome(n int, k int) string {
	ans, pow := make([]byte, n), make([]int, n)
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = pow[i-1] * 10 % k
	}
	m := (n + 1) / 2
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, k)
	}

	var dfs func(i, rem int) bool
	dfs = func(i, rem int) bool {
		if i == m {
			return rem == 0
		}
		if vis[i][rem] {
			return false
		}
		vis[i][rem] = true

		base := pow[i]
		if n-1-i != i {
			base = (base + pow[n-1-i]) % k
		}

		for d := 9; d >= 0; d-- {
			if dfs(i+1, (rem+base*d)%k) {
				ch := byte('0' + d)
				ans[i], ans[n-1-i] = ch, ch
				return true
			}
		}
		return false
	}
	dfs(0, 0)
	return string(ans)
}
