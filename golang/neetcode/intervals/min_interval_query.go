package intervals

import (
	"container/heap"
	"slices"
)

type Heap[T any] struct {
	a    []T
	less func(a, b T) bool
}

func (h *Heap[T]) Len() int           { return len(h.a) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.a[i], h.a[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.a[i], h.a[j] = h.a[j], h.a[i] }
func (h *Heap[T]) Push(v any)         { h.a = append(h.a, v.(T)) }
func (h *Heap[T]) Top() T             { return h.a[0] }
func (h *Heap[T]) Pop() any {
	n := len(h.a)
	v := h.a[n-1]
	h.a = h.a[:n-1]
	return v
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{nil, less}
	heap.Init(h)
	return h
}

func minInterval(intervals [][]int, queries []int) []int {
	n, m := len(intervals), len(queries)
	order, ans := make([]int, m), make([]int, m)
	for i, _ := range order {
		order[i] = i
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	slices.SortFunc(order, func(a, b int) int {
		return queries[a] - queries[b]
	})

	minHeap := NewHeap(func(a, b []int) bool {
		return (a[1] - a[0]) < (b[1] - b[0])
	})

	j := 0
	for _, i := range order {
		q := queries[i]
		for ; j < n && intervals[j][0] <= q; j++ {
			heap.Push(minHeap, intervals[j])
		}
		for minHeap.Len() > 0 && minHeap.Top()[1] < q {
			heap.Pop(minHeap)
		}
		ans[i] = -1
		if minHeap.Len() > 0 {
			mn := minHeap.Top()
			ans[i] = mn[1] - mn[0] + 1
		}
	}
	return ans
}
