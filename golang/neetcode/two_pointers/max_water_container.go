package two_pointers

func maxArea(h []int) int {
	l, r := 0, len(h)-1
	mx := 0
	for l < r {
		mx = max(mx, min(h[l], h[r])*(r-l))
		if h[l] <= h[r] { // l done its max, safe to eliminate
			l++
		} else {
			r--
		}
	}
	return mx
}
