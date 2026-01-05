package linked_list

type MyCircularQueue struct {
	hd, tail *ListNode
	cnt, cap int
}

func ConstructorCQ(k int) MyCircularQueue {
	hd := &ListNode{}
	hd.Next = hd
	for i := 0; i < k-1; i++ {
		hd.Next = &ListNode{Next: hd.Next}
	}
	return MyCircularQueue{hd: hd.Next, tail: hd, cap: k}
}

func (ci *MyCircularQueue) EnQueue(value int) bool {
	if ci.IsFull() {
		return false
	}
	ci.tail = ci.tail.Next
	ci.tail.Val = value
	ci.cnt++
	return true
}

func (ci *MyCircularQueue) DeQueue() bool {
	if ci.IsEmpty() {
		return false
	}
	ci.hd = ci.hd.Next
	ci.cnt--
	return true
}

func (ci *MyCircularQueue) Front() int {
	if ci.IsEmpty() {
		return -1
	}
	return ci.hd.Val
}

func (ci *MyCircularQueue) Rear() int {
	if ci.IsEmpty() {
		return -1
	}
	return ci.tail.Val
}

func (ci *MyCircularQueue) IsEmpty() bool {
	return ci.cnt == 0
}

func (ci *MyCircularQueue) IsFull() bool {
	return ci.cnt >= ci.cap
}
