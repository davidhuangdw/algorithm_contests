package sliding_window

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func checkInclusion(s1 string, s2 string) bool {
	cnt, diff := make(map[byte]int), 0
	n := len(s1)
	for _, c := range []byte(s1) {
		cnt[c]++
		diff++
	}
	add := func(c byte, k int) {
		// diff := sum(Abs(cnt[ch])), diff==0 <=> s1==s2
		diff -= Abs(cnt[c])
		cnt[c] += k
		diff += Abs(cnt[c])
	}

	for i, c := range []byte(s2) {
		add(c, -1)
		if i-n >= 0 {
			add(s2[i-n], 1)
		}
		if diff == 0 {
			return true
		}
	}
	return diff == 0
}
