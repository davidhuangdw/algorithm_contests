package main

import (
	"math/bits"
	"slices"
)

func minimumOR(grid [][]int) int {
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	kb := bits.Len(uint(mx))

	mask := (1 << kb) - 1
	for bit := 1 << (kb - 1); bit > 0; bit >>= 1 {
		find := func(row []int, bit int) bool {
			for _, v := range row {
				if v&bit == 0 {
					return true
				}
			}
			return false
		}
		allRowsFound := func() bool {
			for _, row := range grid {
				if !find(row, bit) {
					return false
				}
			}
			return true
		}
		if allRowsFound() { // could be zero-bit
			mask -= bit
			for i := range grid {
				var noBit []int
				for _, v := range grid[i] {
					if v&bit == 0 { // valid
						noBit = append(noBit, v)
					}
				}
				grid[i] = noBit
			}
		}
	}
	return mask
}
