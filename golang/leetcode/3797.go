package main

import "math"

func numberOfRoutes(grid []string, d int) int {
	M, n, m := int(1e9+7), len(grid), len(grid[0])
	d1 := int(math.Sqrt(float64(d*d) - 1))

	f, g := make([]int, m+1), make([]int, m+1)
	for i := n - 1; i >= 0; i-- {
		// g
		for j := 0; j < m; j++ {
			g[j+1] = g[j]
			if grid[i][j] != '.' {
				continue
			}
			if i == n-1 {
				g[j+1] += 1
			} else {
				g[j+1] += f[min(j+d1+1, m)] - f[max(0, j-d1)]
			}
			g[j+1] %= M
		}
		// f
		for j := 0; j < m; j++ {
			f[j+1] = f[j]
			if grid[i][j] != '.' {
				continue
			}
			f[j+1] += g[min(j+d+1, m)] - g[max(0, j-d)]
			f[j+1] %= M
		}
	}
	return (f[m] + M) % M
}
