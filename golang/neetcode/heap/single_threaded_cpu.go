package heap

import (
	"container/heap"
	"slices"
)

func getOrder(tasks [][]int) []int {
	cand := NewHeap(func(i, j int) bool {
		if tasks[i][1] != tasks[j][1] {
			return tasks[i][1] < tasks[j][1]
		}
		return i < j
	})
	time, n := 0, len(tasks)

	sorted := make([]int, n) // sort by start
	for i := range sorted {
		sorted[i] = i
	}
	slices.SortFunc(sorted, func(i, j int) int {
		return tasks[i][0] - tasks[j][0]
	})

	ans, i := make([]int, 0), 0
	for i < n || cand.Len() > 0 {
		// collect candidates
		for i < n && tasks[sorted[i]][0] <= time {
			heap.Push(cand, sorted[i])
			i++
		}
		// pick
		if cand.Len() == 0 {
			time = tasks[sorted[i]][0]
			continue
		}
		k := heap.Pop(cand).(int)
		ans = append(ans, k)
		time += tasks[k][1]
	}
	return ans
}