package main

func countVisiblePeople(n int, pos int, k int) int {
	M := int(1e9 + 7)
	pow := func(a, b int) int {
		res := 1
		for b > 0 {
			if b&1 > 0 {
				res = res * a % M
			}
			a, b = a*a%M, b/2
		}
		return res
	}

	comb := func(n, k int) int {
		if k > n-k {
			k = n - k
		}
		a, b := 1, 1
		for i := 1; i <= k; i++ {
			a = a * (n - i + 1) % M
			b = b * i % M
		}
		return a * pow(b, M-2) % M
	}

	return comb(n-1, k) * 2 % M
}
