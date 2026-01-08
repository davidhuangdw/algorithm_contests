package backtrace

import (
	"fmt"
	"slices"
)

func makesquare(matchsticks []int) bool {
	sum, n := 0, len(matchsticks)
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sum /= 4
	slices.SortFunc(matchsticks, func(a, b int) int {
		return b - a
	})
	vis := make([]bool, len(matchsticks))
	var edges [4][]int
	var tryEdge func(dep, st, preSum int) bool
	tryEdge = func(dep, st, preSum int) bool {
		if dep == 3 {
			fmt.Printf("%#v", edges)
			return true
		}
		if preSum == sum {
			return tryEdge(dep+1, 0, 0)
		}
		if preSum > sum || st >= n {
			return false
		}
		for i := st; i < n; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			edges[dep] = append(edges[dep], matchsticks[i])
			if tryEdge(dep, i+1, preSum+matchsticks[i]) {
				return true
			}
			vis[i] = false
			edges[dep] = edges[dep][:len(edges[dep])-1]
			if preSum == 0 { // first item just directly choose first
				break
			}
		}
		return false
	}
	return tryEdge(0, 0, 0)
}
