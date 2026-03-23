package main

import "container/heap"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// toList converts a slice of integers to a linked list.
func toList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

// fromList converts a linked list back to a slice of integers.
func fromList(head *ListNode) []int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
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
