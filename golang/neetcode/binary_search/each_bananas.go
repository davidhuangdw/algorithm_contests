package binary_search

func minEatingSpeed(piles []int, h int) int {
	succ := func(k int) bool {
		s := 0
		for _, v := range piles {
			s += (v + k - 1) / k
		}
		return s <= h
	}
	l, r := 1, int(2e9) // fail, succ
	for l <= r {
		md := (l + r) / 2
		if succ(md) {
			r = md - 1
		} else {
			l = md + 1
		}
	}
	return l
}
