package main

func minLength(s string, numOps int) int {
	n, cnt := len(s), 0
	for i, c := range []byte(s) {
		v := int(c - '0')
		if v == i%2 {
			cnt++
		}
	}
	cnt = min(cnt, n-cnt)
	if cnt <= numOps {
		return 1
	}
	var seg []int
	for i, j := 0, 0; i < n; i = j {
		for j = i + 1; j < n && s[j] == s[i]; j++ {
		}
		seg = append(seg, j-i)
	}
	l, r := 2, n
	for l <= r {
		md, rem := (l+r)>>1, numOps
		for i := 0; i < len(seg) && rem >= 0; i++ {
			rem -= seg[i] / (md + 1)
		}
		if rem < 0 {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return l
}
