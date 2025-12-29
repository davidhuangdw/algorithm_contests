package dp

import "strconv"

func numDecodings(s string) int {
	valid := func(sub string) bool {
		v, _ := strconv.Atoi(sub)
		return sub[0] != '0' && v <= 26
	}
	a, b := 0, 1
	for i := 0; i < len(s) && a+b > 0; i++ {
		cur := 0
		if valid(s[i : i+1]) {
			cur += b
		}
		if i-1 >= 0 && valid(s[i-1:i+1]) {
			cur += a
		}
		a, b = b, cur
	}
	return b
}
