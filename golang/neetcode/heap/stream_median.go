package heap

import "container/heap"

type MedianFinder struct {
	maxHeap, minHeap *Heap[int]
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewHeap(func(a, b int) bool { return a > b }),
		minHeap: NewHeap(func(a, b int) bool { return a < b }),
	}
}

func (m *MedianFinder) AddNum(num int) {
	if m.maxHeap.Len() == 0 || num <= m.maxHeap.Top() {
		heap.Push(m.maxHeap, num)
		if m.maxHeap.Len()-m.minHeap.Len() > 1 {
			heap.Push(m.minHeap, heap.Pop(m.maxHeap))
		}
	} else {
		heap.Push(m.minHeap, num)
		if m.minHeap.Len() > m.maxHeap.Len() {
			heap.Push(m.maxHeap, heap.Pop(m.minHeap))
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.maxHeap.Len() == 0 {
		return 0
	}
	if m.maxHeap.Len() > m.minHeap.Len() {
		return float64(m.maxHeap.Top())
	}
	return float64(m.minHeap.Top()+m.maxHeap.Top()) / 2

}
