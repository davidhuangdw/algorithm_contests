package main

import (
	"slices"
)

func longestPalindrome(s string, t string) int {
	a, b := []byte(s), []byte(t)
	slices.Reverse(b)
	calc := func(a, b []byte) int {
		n, res := len(a), 0
		match, dp := make([]int, n+1), make([]int, n+1) // match[i+1] := max len for a[:i]
		for _, c := range b {
			for i := n - 1; i >= 0; i-- {
				if a[i] == c {
					dp[i+1] = dp[i] + 1
				} else {
					dp[i+1] = 0
				}
				match[i+1] = max(match[i+1], dp[i+1])
			}
		}
		//fmt.Println(match)

		expand := func(i, j int) (l, r int) {
			for l, r = i, j; 0 <= l && r < n && a[l] == a[r]; {
				l--
				r++
			}
			return l, r
		}
		for i := 0; i < n; i++ {
			// expand
			l, r := expand(i, i)
			res = max(res, r-1-l+match[l+1]*2)
			l, r = expand(i, i+1)
			res = max(res, r-1-l+match[l+1]*2)
		}
		return res
	}
	return max(calc(a, b), calc(b, a))
}
