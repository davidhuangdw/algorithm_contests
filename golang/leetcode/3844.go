package main

func almostPalindromic(s string) int {
	a, n := []byte(s), len(s)
	var expand func(i, j, diff int) int
	expand = func(i, j, diff int) int {
		for i >= 0 && j < n && a[i] == a[j] {
			i--
			j++
		}
		if diff == 0 {
			return j - i - 1
		}
		mx := j - i - 1
		if i >= 0 {
			mx = max(mx, expand(i-1, j, diff-1))
		}
		if j < n {
			mx = max(mx, expand(i, j+1, diff-1))
		}
		return mx
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, expand(i, i, 1), expand(i, i+1, 1))
	}
	return ans
}
