package heap

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := NewHeap(func(a, b int) bool { return a < b })
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return h.Top()
}
