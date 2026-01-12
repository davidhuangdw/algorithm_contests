package heap

import "container/heap"

func carPooling(trips [][]int, capacity int) bool {
	points := NewHeap(func(a, b []int) bool {
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})
	for _, tri := range trips {
		k, fr, to := tri[0], tri[1], tri[2]
		heap.Push(points, []int{fr, +k})
		heap.Push(points, []int{to, -k})
	}
	sz := 0
	for points.Len() > 0 {
		p := heap.Pop(points).([]int)
		sz += p[1]
		if sz > capacity {
			return false
		}
	}
	return true
}
