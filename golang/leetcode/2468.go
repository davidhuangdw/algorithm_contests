package main

import (
	"fmt"
)

func splitMessage(message string, limit int) []string {
	n := len(message)
	cnt, nd := 0, 0
	for b, i := 1, 0; i < n; b *= 10 {
		nd++
		i -= b - 1 // delta to fix prev
		rem := limit - (3 + nd*2)
		if rem <= 0 {
			return nil
		}
		add := min(((n-i)+rem-1)/rem, b*10-b)
		i += add * rem
		cnt += add
	}
	ans := make([]string, 0)
	i, s := 0, []byte(message)
	for d, b := 1, 1; d <= nd; d, b = d+1, b*10 {
		for k := b; k < min(cnt+1, b*10); k++ {
			j := min(n, i+(limit-(3+nd+d)))
			ans, i = append(ans, fmt.Sprintf("%s<%d/%d>", s[i:j], k, cnt)), j
		}
	}
	return ans
}
