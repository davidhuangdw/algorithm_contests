package main

func balancedString(s string) int {
	n := len(s)
	avg, cnt := n/4, make(map[byte]int)
	for _, c := range []byte(s) {
		cnt[c]++
	}
	more := 0
	for _, k := range cnt {
		more += max(0, k-avg)
	}
	if more <= 1 {
		return more
	}

	ans, r := n, 0
	for l := range []byte(s) {
		for ; r < n && more > 0; r++ {
			if cnt[s[r]] > avg {
				more--
			}
			cnt[s[r]]--
		}
		if more == 0 {
			ans = min(ans, r-l)
		}
		cnt[s[l]]++
		if cnt[s[l]] > avg {
			more++
		}
	}
	return ans
}
