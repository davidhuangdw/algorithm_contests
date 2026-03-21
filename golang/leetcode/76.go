package main

func minWindow(s string, t string) string {
	rem := make(map[byte]int)
	for _, v := range []byte(t) {
		rem[v]++
	}
	fail, n, r := len(rem), len(s), 0
	add := func(v byte, d int) {
		if _, ok := rem[v]; ok {
			rem[v] += d
			if d < 0 && rem[v] == 0 {
				fail--
			}
			if d > 0 && rem[v] == 1 {
				fail++
			}
		}
	}
	var ans []byte
	bs := []byte(s)
	for l, v := range bs {
		for r <= l || (fail > 0 && r < n) {
			add(s[r], -1)
			r++
		}
		if fail == 0 && (len(ans) == 0 || len(ans) > r-l) {
			ans = bs[l:r]
		}
		add(v, 1)
	}
	return string(ans)
}
