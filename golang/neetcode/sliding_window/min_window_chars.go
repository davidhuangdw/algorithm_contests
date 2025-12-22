package sliding_window

func minWindow(ss string, tt string) string {
	s, t := []byte(ss), []byte(tt)
	cnt := make(map[byte]int)
	for _, c := range t {
		cnt[c]--
	}
	neg, i := len(cnt), 0 // neg := how many cnt[ch]<0 (~= ch is not covered yet)

	add := func(c byte, k int) {
		//if cnt[c] >= 0 && cnt[c]+k < 0 {
		//	neg++
		//}
		if cnt[c] < 0 && cnt[c]+k >= 0 {
			neg--
		}
		cnt[c] += k
	}

	minLen := len(ss) + 1
	var ans []byte
	for j, c := range s {
		add(c, 1)
		for i <= j && cnt[s[i]] > 0 {
			add(s[i], -1)
			i++
		}
		if neg == 0 && j+1-i < minLen {
			minLen = j + 1 - i
			ans = s[i : j+1]
		}
	}
	return string(ans)
}
