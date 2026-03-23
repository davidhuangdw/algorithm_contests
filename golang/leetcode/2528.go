package main

func maxPower(stations []int, rad int, k int) int64 {
	n := len(stations)
	base := make([]int, n)
	for i, v := range stations {
		base[max(0, i-rad)] += v
		if i+rad+1 < n {
			base[i+rad+1] = -v
		}
	}
	check := func(want int) bool {
		extra := make([]int, n)
		rem, pre := k, 0
		for i, vb := range base {
			pre += vb + extra[i]
			if pre < want {
				d := want - pre
				rem -= d
				if rem < 0 {
					return false
				}
				pre = want
				if i+rad*2+1 < n {
					extra[i+rad*2+1] += -d
				}
			}
		}
		return true
	}
	l, r := 0, int(2e10)
	for l <= r {
		md := (l + r) >> 1
		if check(md) {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return int64(r)
}
