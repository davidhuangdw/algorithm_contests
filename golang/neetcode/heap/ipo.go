package heap

import (
	"container/heap"
	"slices"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n, cand := len(capital), NewHeap(func(a, b int) bool {
		return a > b
	})

	byCap := make([]int, n)
	for i := range byCap {
		byCap[i] = i
	}
	slices.SortFunc(byCap, func(i, j int) int {
		return capital[i] - capital[j]
	})

	sum := w
	for t := 0; t < k; t++ {
		for len(byCap) > 0 && capital[byCap[0]] <= sum {
			heap.Push(cand, profits[byCap[0]])
			byCap = byCap[1:]
		}
		if cand.Len() == 0 {
			break
		}
		sum += heap.Pop(cand).(int)
	}
	return sum
}
