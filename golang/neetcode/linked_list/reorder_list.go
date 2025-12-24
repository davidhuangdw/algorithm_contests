package linked_list

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// cut
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	p := head

	// reverse cur into q
	var q *ListNode
	for cur != nil {
		nxt := cur.Next
		cur.Next = q
		q = cur
		cur = nxt
	}

	// inter merge
	hd := &ListNode{}
	cur = hd
	for p != nil && q != nil {
		cur.Next = p
		p = p.Next
		cur.Next.Next = q
		q = q.Next
		cur = cur.Next.Next
	}
	cur.Next = p
}
