package main

func lengthAfterTransformations(s string, t int, nums []int) int {
	type m26 [26][26]int
	M := int(1e9 + 7)
	mul := func(a, b m26) m26 {
		var res m26
		for i := range 26 {
			for j := range 26 {
				s := &res[i][j]
				for k := range 26 {
					*s = (*s + a[i][k]*b[k][j]) % M
				}
			}
		}
		return res
	}
	pow := func(a m26, k int) m26 {
		var res m26
		for i := range 26 {
			res[i][i] = 1
		}
		for p := a; k > 0; k /= 2 {
			if k%2 == 1 {
				res = mul(res, p)
			}
			p = mul(p, p)
		}
		return res
	}
	var mt m26
	for i, v := range nums {
		for d := range v {
			mt[i][(i+1+d)%26] = 1
		}
	}
	mt = pow(mt, t)
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a']++
	}
	ans := 0
	for i, c := range cnt {
		for j := range 26 {
			ans += mt[i][j] * c % M
		}
	}
	return ans % M
}
