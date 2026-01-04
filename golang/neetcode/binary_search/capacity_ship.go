package binary_search

func shipWithinDays(weights []int, days int) int {
	mx, sum := 0, 0
	for _, v := range weights {
		mx = max(mx, v)
		sum += v
	}
	l, r := mx, sum
	for l <= r {
		m, pre, k := (l+r)/2, 0, 1
		for _, v := range weights {
			pre += v
			if pre > m {
				pre, k = v, k+1
			}
		}
		if k <= days {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}