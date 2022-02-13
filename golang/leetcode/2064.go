package leetcode

func minimizedMaximum(n int, quantities []int) int {
	poss := func(mx int) bool{
		need := 0
		for _, q := range quantities {
			need += (q+mx-1)/mx
			if need > n {
				return false
			}
		}
		return true
	}
	lo, hi := 1, 100_000
	for lo <= hi {
		m := (lo + hi) / 2
		if poss(m) {
			hi = m-1
		}else {
			lo = m+1
		}
	}
	return lo
}
