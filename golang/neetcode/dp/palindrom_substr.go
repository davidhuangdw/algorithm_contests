package dp

func longestPalindrome1(s string) string {
	n, mx := len(s), s[:0]
	for i := 0; i < n; i++ {
		for _, start := range [][]int{{i, i}, {i, i + 1}} { // odd & even
			l, r := start[0], start[1]
			for 0 <= l && r < n && s[l] == s[r] {
				l--
				r++
			}
			if len(mx) < r-1-l {
				mx = s[l+1 : r]
			}
		}
	}
	return mx
}

func longestPalindrome(s string) string {
	mx := s[:0]
	for i, l := range manacher(s) {
		cur := s[i/2-(l-1)/2 : i/2+l/2+1] // left := i/2 - (l-1)/2, less cut for even case
		if len(cur) > len(mx) {
			mx = cur
		}
	}
	return mx
}
