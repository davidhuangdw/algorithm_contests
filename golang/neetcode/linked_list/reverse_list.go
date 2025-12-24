package linked_list

func reverseList(head *ListNode) *ListNode {
	var hd *ListNode
	for p := head; p != nil; {
		nxt := p.Next
		p.Next = hd
		hd = p
		p = nxt
	}
	return hd
}
