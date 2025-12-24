package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// forward n step
	ed := head
	for ; n > 0 && ed != nil; n-- {
		ed = ed.Next
	}
	if n > 0 {
		return head
	}

	// locate end
	hd := &ListNode{Next: head}
	prev := hd
	for ed != nil {
		prev = prev.Next
		ed = ed.Next
	}

	// remove
	prev.Next = prev.Next.Next
	return hd.Next
}
