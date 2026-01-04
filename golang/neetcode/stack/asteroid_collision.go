package stack

func asteroidCollision(a []int) []int {
	j := -1
	for _, v := range a {
		for v < 0 && j >= 0 && a[j] > 0 {
			if a[j] < -v {
				j--
				continue
			}
			if a[j] == -v {
				j--
			}
			v = 0
		}
		if v != 0 {
			j++
			a[j] = v
		}
	}
	return a[:j+1]
}
