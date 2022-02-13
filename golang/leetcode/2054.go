package leetcode

import "sort"

type Event struct {
	st, ed, val int
}

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool { return events[i][0] > events[j][0] })
	type maxVal struct {
		st, max int
	}
	gt := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	var maxSince []maxVal
	pop := func() {
		maxSince = maxSince[:len(maxSince)-1]
	}
	push := func(x maxVal) {
		maxSince = append(maxSince, x)
	}
	last := func() maxVal {
		return maxSince[len(maxSince)-1]
	}
	empty := func() bool {
		return len(maxSince) == 0
	}

	for i := 0; i < n; {
		st := events[i][0]
		max := events[i][2]
		ni := n
		for j, e := range events[i+1:] {
			if e[0] != st {
				ni = j + i + 1
				break
			} else {
				max = gt(max, e[2])
			}
		}
		if empty() || last().max < max {
			push(maxVal{st, max})
		}
		i = ni
	}

	max := 0
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	for _, e := range events {
		v := e[2]
		ed := e[1]
		for !empty() && last().st <= ed {
			pop()
		}
		if !empty() {
			v += last().max
		}
		max = gt(max, v)
	}
	return max
}
