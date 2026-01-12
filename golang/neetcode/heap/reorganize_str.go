package heap

func reorganizeString(s string) string {
	// strategy: continuously put the same char for i+2
	cnt, mx, n := make(map[byte]int), byte(0), len(s)
	for _, ch := range []byte(s) {
		cnt[ch]++
		if cnt[ch] > cnt[mx] {
			mx = ch
		}
	}
	if cnt[mx]*2-1 > len(s) {
		return ""
	}

	i, res := 0, make([]byte, n)
	nxt := func(i int) int {
		if i+2 < n {
			return i + 2
		}
		return 1
	}
	for cnt[mx] > 0 {
		cnt[mx]--
		res[i] = mx
		i = nxt(i)
	}
	delete(cnt, mx)

	for b, c := range cnt {
		for ; c > 0; c-- {
			res[i] = b
			i = nxt(i)
		}
	}
	return string(res)
}
