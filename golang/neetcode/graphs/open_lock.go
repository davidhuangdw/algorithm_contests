package graphs

import (
	"strconv"
)

func openLock(deadends []string, target string) int {
	dead := make(map[int]bool)
	for _, d := range deadends {
		num, _ := strconv.Atoi(d)
		dead[num] = true
	}
	tarNum, _ := strconv.Atoi(target)
	if tarNum == 0 {
		return 0
	}

	var que [][]int            // num, steps
	qued := make(map[int]bool) // nums
	if !dead[0] {
		que = append(que, []int{0, 0})
		qued[0] = true
	}

	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		num, steps := p[0], p[1]
		// for next transit
		for b, t := 1, 0; t < 4; t++ {
			dig := num / b % 10
			for _, nxtDig := range []int{(dig + 9) % 10, (dig + 1) % 10} {
				nxtNum := num + (nxtDig-dig)*b
				if qued[nxtNum] || dead[nxtNum] {
					continue
				}
				if nxtNum == tarNum {
					return steps + 1
				}
				qued[nxtNum] = true
				que = append(que, []int{nxtNum, steps + 1})
			}
			b *= 10
		}
	}
	return -1
}
