package main

import (
	"slices"
)

func countFancy(l int64, r int64) int64 {
	toDigits := func(v int64) []int {
		var d []int
		for ; v > 0; v /= 10 {
			d = append(d, int(v%10))
		}
		slices.Reverse(d)
		return d
	}

	isFancy := func(n int) bool {
		s := toDigits(int64(n))
		inc, dec := true, true
		for i := 1; i < len(s); i++ {
			if s[i] <= s[i-1] {
				inc = false
			}
			if s[i] >= s[i-1] {
				dec = false
			}
		}
		return inc || dec
	}

	var isGood [9*15 + 1]bool
	for i := range isGood {
		isGood[i] = isFancy(i)
	}

	countGood := func(n int64) int64 {
		if n < 0 {
			return 0
		}
		type state struct {
			i, pre, sum, trend int
			isLimit, begin     bool
		}
		memo, digits := make(map[state]int64), toDigits(n)
		// trend: 0 init, 1 trend, 2 dec, 3 fail
		var dp func(i, pre, sum, trend int, isLimit, begin bool) int64
		dp = func(i, pre, sum, trend int, isLimit, begin bool) int64 {
			if i == len(digits) {
				if !begin {
					return 0
				}
				if trend != 3 || isGood[sum] {
					return 1
				}
				return 0
			}

			st := state{i, pre, sum, trend, isLimit, begin}
			if _, ok := memo[st]; !ok {
				var res int64
				fr, to := 0, 9
				if !begin {
					res += dp(i+1, pre, sum, trend, false, false) // bug: isLimit:=false
					fr = 1
				}

				if isLimit {
					to = digits[i]
				}

				for d := fr; d <= to; d++ {
					tr := trend
					switch {
					case !begin:
						tr = 0
					case pre < d && (trend == 0 || trend == 1):
						tr = 1
					case pre > d && (trend == 0 || trend == 2):
						tr = 2
					default:
						tr = 3
					}
					res += dp(i+1, d, sum+d, tr, isLimit && d == digits[i], true)
				}

				memo[st] = res
			}
			return memo[st]
		}

		return dp(0, -1, 0, 0, true, false)
	}

	return countGood(r) - countGood(l-1)
}
