package binary_search

type MountainArray interface {
	get(index int) int
	length() int
}

func findInMountainArray(target int, a MountainArray) int {
	n := a.length()

	l, r := 0, n-2
	for l <= r {
		m := (l + r) / 2
		if a.get(m) < a.get(m+1) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	top := l
	bsearch := func(l, r int, up bool) int {
		for l <= r {
			m := (l + r) / 2
			mv := a.get(m)
			if mv == target {
				return m
			}
			if (up && target < mv) || (!up && target > mv) {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return -1
	}
	i := bsearch(0, top, true)
	if i >= 0 {
		return i
	}
	return bsearch(top, n-1, false)
}
