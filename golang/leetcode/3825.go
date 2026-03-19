package main

func longestSubsequence(nums []int) int {
	lis := func(a []int) int {
		mx := 0
		var inc []int
		for _, v := range a {
			l, r := 0, len(inc)-1
			for l <= r {
				md := (l + r) / 2
				if v > inc[md] {
					l = md + 1
				} else {
					r = md - 1
				}
			}
			if l < len(inc) {
				inc[l] = v
			} else {
				inc = append(inc, v)
			}
			mx = max(mx, len(inc))
		}
		return mx
	}
	ans := 0
	for b := 1; b <= 1<<30; b <<= 1 {
		var a []int
		for _, v := range nums {
			if v&b > 0 {
				a = append(a, v)
			}
		}
		ans = max(ans, lis(a))
	}
	return ans
}
