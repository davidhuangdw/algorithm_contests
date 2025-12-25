package linked_list

import "github.com/emirpasic/gods/queues/priorityqueue"

func mergeKLists1(lists []*ListNode) *ListNode {
	hd := new(ListNode)
	tail := hd

	for {
		minI := -1
		for i, nd := range lists {
			if nd == nil {
				continue
			}
			if minI < 0 || nd.Val < lists[minI].Val {
				minI = i
			}
		}
		if minI < 0 {
			break
		}
		nd := lists[minI]
		lists[minI] = nd.Next
		tail.Next = nd
		tail = nd
	}
	return hd.Next
}

func mergeKLists2(lists []*ListNode) *ListNode { // O(logk * n) by priority-queue
	que := priorityqueue.NewWith(func(a, b any) int {
		return a.(*ListNode).Val - b.(*ListNode).Val
	})
	for _, nd := range lists {
		if nd != nil {
			que.Enqueue(nd)
		}
	}

	hd := new(ListNode)
	tail := hd
	for !que.Empty() {
		v, _ := que.Dequeue()
		nd := v.(*ListNode)
		tail.Next = nd
		tail = tail.Next
		if nd.Next != nil {
			que.Enqueue(nd.Next)
		}
	}
	return hd.Next
}

func mergeKLists(lists []*ListNode) *ListNode { // O(logk * n) by divide-conquer
	var merge func(l, r int) *ListNode
	merge = func(l, r int) *ListNode {
		if l > r {
			return nil
		}
		if l == r {
			return lists[l]
		}
		md := (l + r) / 2
		a := merge(l, md)
		b := merge(md+1, r)

		hd := new(ListNode)
		tail := hd
		for a != nil && b != nil {
			if a.Val <= b.Val {
				tail.Next = a
				a = a.Next
			} else {
				tail.Next = b
				b = b.Next
			}
			tail = tail.Next
		}
		tail.Next = a
		if a == nil {
			tail.Next = b
		}
		return hd.Next
	}
	return merge(0, len(lists)-1)
}
