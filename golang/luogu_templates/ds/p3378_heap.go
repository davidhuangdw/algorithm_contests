package ds

import (
	"cmp"
	"fmt"
	"io"
)

type Heap[T cmp.Ordered] struct {
	data []T
}

func NewHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{data: make([]T, 1)}
}

func (h *Heap[T]) siftUp(i int) {
	for p := i / 2; p >= 1 && !cmp.Less(h.data[p], h.data[i]); i, p = p, p/2 {
		h.data[i], h.data[p] = h.data[p], h.data[i]
	}
}
func (h *Heap[T]) siftDown(i int) {
	for ch := i * 2; ch < len(h.data); i, ch = ch, ch*2 {
		if ch+1 < len(h.data) && cmp.Less(h.data[ch+1], h.data[ch]) {
			ch++
		}
		if cmp.Less(h.data[i], h.data[ch]) {
			break
		}
		h.data[i], h.data[ch] = h.data[ch], h.data[i]
	}
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.siftUp(len(h.data) - 1)
}
func (h *Heap[T]) Pop() T {
	top := h.data[1]
	h.data[1], h.data = h.data[len(h.data)-1], h.data[:len(h.data)-1]
	h.siftDown(1)
	return top
}
func (h *Heap[T]) Top() T {
	return h.data[1]
}
func (h *Heap[T]) Len() int {
	return len(h.data) - 1
}

func P3378_heap(IN io.Reader, OUT io.Writer) {
	var k int
	fmt.Fscan(IN, &k)
	h := NewHeap[int]()
	for ; k > 0; k-- {
		var op, v int
		fmt.Fscan(IN, &op)
		switch op {
		case 1:
			fmt.Fscan(IN, &v)
			h.Push(v)
		case 2:
			fmt.Fprintln(OUT, h.Top())
		case 3:
			h.Pop()
		}
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3378_heap(IN, OUT)
//}
