package sliding_window

func lengthOfLongestSubstring(ss string) int {
	s := []rune(ss)
	prev, l, mx := make(map[rune]int), 0, 0
	for r, c := range s {
		if p, ok := prev[c]; ok {
			l = max(l, p+1)
		}
		prev[c] = r
		mx = max(mx, r+1-l)
	}
	return mx
}
