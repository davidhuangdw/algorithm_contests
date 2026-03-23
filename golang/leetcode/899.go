package main

import "slices"

func orderlyQueue(s string, k int) string {
	if k >= 2 {
		b := []byte(s)
		slices.Sort(b)
		return string(b)
	}
	i, j, n, k := 0, 1, len(s), 0
	for i < n && j < n && k < n {
		a, b := s[(i+k)%n], s[(j+k)%n]
		if a == b {
			k++
			continue
		}
		if a < b {
			j += k + 1
		} else {
			i += k + 1
		}
		k = 0
		if i == j {
			j++
		}
	}
	i = min(i, j)
	return s[i:] + s[:i]
}
