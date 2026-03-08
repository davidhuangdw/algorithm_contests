package main

import "slices"

// q3
func minOperations(s string) int {
	n := len(s)
	if n == 2 && s[0] > s[1] {
		return -1
	}
	if slices.IsSorted([]byte(s)) {
		return 0
	}
	mn, mx := s[0], s[0]
	for _, c := range []byte(s) {
		mn = min(mn, c)
		mx = max(mx, c)
	}
	if s[0] == mn || s[n-1] == mx {
		return 1
	}
	count := func(c byte) int {
		cnt := 0
		for _, v := range []byte(s) {
			if v == c {
				cnt++
			}
		}
		return cnt
	}
	mnCnt, mxCnt := count(mn), count(mx)
	if mnCnt == 1 && mxCnt == 1 && s[0] == mx && s[n-1] == mn {
		return 3
	}
	return 2
}
