package main

import (
	"slices"
)

func maximumCoins(coins [][]int, k int) int64 {
	ans, n := 0, len(coins)
	for t := range 2 {
		slices.SortFunc(coins, func(a, b []int) int {
			return a[0] - b[0]
		})
		var ri, sum int
		for _, l := range coins {
			ed := l[0] + k - 1
			for ri < n && coins[ri][0] <= ed {
				sum += (coins[ri][1] + 1 - coins[ri][0]) * coins[ri][2]
				ri++
			}
			r := coins[ri-1]
			ans = max(ans, sum-max(0, r[1]-ed)*r[2])
			sum -= (l[1] + 1 - l[0]) * l[2]
		}
		if t == 0 { //flip & retry
			for _, c := range coins {
				c[0], c[1] = -c[1], -c[0]
			}
		}
	}
	return int64(ans)
}
