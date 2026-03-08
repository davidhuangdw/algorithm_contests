package main

func maximumXor(s string, t string) string {
	var cnt [2]int
	for _, c := range []byte(t) {
		cnt[c-'0']++
	}
	ans := make([]byte, len(s))
	for i, c := range []byte(s) {
		d, p := int(c-'0'), 0
		if cnt[d^1] > 0 {
			p = d ^ 1
		} else {
			p = d
		}
		cnt[p]--
		ans[i] = byte('0' + d ^ p)
	}
	return string(ans)
}
