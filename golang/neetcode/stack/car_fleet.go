package stack

import (
	"slices"
)

func carFleet(target int, position []int, speed []int) int {
	n := len(speed)
	if n == 0 {
		return 0
	}
	ord := make([]int, n)
	for i, _ := range ord {
		ord[i] = i
	}
	slices.SortFunc(ord, func(i, j int) int { // desc
		if position[i] == position[j] {
			return j - i
		}
		return position[j] - position[i]
	})
	cnt, j := 1, ord[0]
	for _, i := range ord[1:] {
		if (target-position[i])*speed[j] <= (target-position[j])*speed[i] {
			continue
		}
		cnt++
		j = i
	}
	return cnt
}
