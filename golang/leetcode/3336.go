package main

import "slices"

func subsequencePairCount(nums []int) int {
	U, M, n := slices.Max(nums)+1, int(1e9+7), len(nums)
	gcd := make([][]int, U)
	for a := range gcd {
		gcd[a] = make([]int, U)
		for b := range gcd[a] {
			x, y := a, b
			for x > 0 {
				x, y = y%x, x
			}
			gcd[a][b] = y
		}
	}
	var mm [2][][]int
	for i := range 2 {
		mm[i] = make([][]int, U)
		for l := range mm[i] {
			mm[i][l] = make([]int, U)
		}
	}
	mm[0][0][0] = 1

	for i, x := range nums {
		cur, nxt := mm[i&1], mm[(i+1)&1]
		for l := range U {
			for r := range U {
				nxt[l][r] = cur[l][r]
			}
		}
		for l := range U {
			for r := range U {
				if cur[l][r] > 0 {
					nxt[gcd[l][x]][r] = (nxt[gcd[l][x]][r] + cur[l][r]) % M
					nxt[l][gcd[r][x]] = (nxt[l][gcd[r][x]] + cur[l][r]) % M
				}
			}
		}
	}
	ans := 0
	for x := 1; x < U; x++ {
		ans = (ans + mm[n&1][x][x]) % M
	}
	return ans
}
