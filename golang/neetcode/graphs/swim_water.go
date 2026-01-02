package graphs

import "container/heap"

func swimInWater(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	done := make([][]bool, n)
	for i, _ := range done {
		done[i] = make([]bool, m)
	}

	minHeap := NewHeap(func(a, b []int) bool {
		return a[2] < b[2]
	})
	heap.Push(minHeap, []int{0, 0, grid[0][0]})
	for minHeap.Len() > 0 {
		v := heap.Pop(minHeap).([]int)
		i, j, d := v[0], v[1], v[2]
		if done[i][j] {
			continue
		}
		done[i][j] = true
		if i == n-1 && j == m-1 {
			return d
		}
		for di, dj, t := 0, 1, 0; t < 4; t++ {
			ni, nj := i+di, j+dj
			if 0 <= ni && ni < n && 0 <= nj && nj < m && !done[ni][nj] {
				heap.Push(minHeap, []int{ni, nj, max(d, grid[ni][nj])})
			}
			di, dj = dj, -di
		}
	}
	return -1
}

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
