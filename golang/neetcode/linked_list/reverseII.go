package linked_list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	hd := ListNode{Next: head}
	pre := &hd
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	r, tail := pre.Next, pre.Next
	for i := 0; r != nil && i < right+1-left; i++ {
		r, pre.Next, r.Next = r.Next, r, pre.Next
	}
	if tail != nil {
		tail.Next = r
	}
	return hd.Next
}
