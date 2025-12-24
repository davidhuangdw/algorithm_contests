package linked_list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	hd := &ListNode{}
	tail := hd

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	tail.Next = l1
	if l2 != nil {
		tail.Next = l2
	}
	return hd.Next
}
