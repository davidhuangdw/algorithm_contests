package main

func smallestPalindrome(s string, k int) string {
	n, hf, bs := len(s), len(s)/2, []byte(s)
	var cnt [26]int
	for _, ch := range bs[:hf] {
		cnt[int(ch-'a')]++
	}
	comb := func(n, m int) int {
		res := 1
		m = min(n-m, m) // trick to avoid overflow
		for i := 1; i <= m; i++ {
			res = res * n / i
			if res >= k {
				return res
			}
			n--
		}
		return res
	}
	remain := func(sz int) int {
		res := 1
		for _, c := range cnt {
			res *= comb(sz, c)
			if res >= k {
				return res
			}
			sz -= c
		}
		return res
	}
	if remain(hf) < k {
		return ""
	}
outer:
	for i := 0; i < hf; i++ {
		for v := range cnt {
			if cnt[v] > 0 {
				cnt[v]--
				r := remain(hf - i - 1)
				if r >= k {
					bs[i] = byte('a' + v)
					continue outer
				} else {
					cnt[v]++
					k -= r
				}
			}
		}
		return ""
	}

	for i := 0; i < hf; i++ {
		bs[n-1-i] = bs[i]
	}
	return string(bs)
}
