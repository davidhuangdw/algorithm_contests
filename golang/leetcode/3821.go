package main

func nthSmallest(n int64, k int) int64 {
	comb := func(m, j int) int {
		res := 1
		for i := 1; i <= j; i++ {
			res *= m + 1 - i
			res /= i
		}
		return res
	}

	nn := int(n)
	var ans int64
	for b := 50; b >= 0; b-- {
		c := comb(b, k)
		if nn > c {
			nn -= c
			k--
			ans |= 1 << b
		}
	}
	return ans
}
