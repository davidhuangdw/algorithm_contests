package main

import "math"

func countGoodIntegers(n int, k int) int64 {
	goods := make(map[[10]int]bool)

	// gen
	fr, to := int(math.Pow10((n-1)/2)), int(math.Pow10((n+1)/2))
	for v := fr; v < to; v++ {
		full, x := v, v
		var dig [10]int
		for i := 0; i < (n+1)/2; i++ {
			d := x % 10
			if i == 0 && n%2 == 1 {
				dig[d] += 1
			} else {
				dig[d] += 2
				full = full*10 + d
			}
			x /= 10
		}
		if full%k == 0 {
			goods[dig] = true
		}
	}

	comb := func(dig [10]int) int {
		a, b, pre := 1, 1, 0
		for _, c := range dig {
			for i := 1; i <= c; i++ {
				pre++
				a *= pre
				b *= i
			}
		}
		return a / b
	}
	// count
	ans := 0
	for dig := range goods {
		ans += comb(dig)
		if dig[0] > 0 {
			dig[0]--
			ans -= comb(dig)
		}
	}
	return int64(ans)
}
