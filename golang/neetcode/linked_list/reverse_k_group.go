package linked_list

func reverseKGroup(head *ListNode, k int) *ListNode {
	hd := new(ListNode)
	prev, tail := hd, hd
	for head != nil {
		tail = head
		i := 0
		for ; i < k && head != nil; i++ {
			prev.Next, head.Next, head = head, prev.Next, head.Next
		}
		if i < k { // restore order of last not-full group
			tail, prev.Next = prev.Next, nil
			for tail != nil {
				tail.Next, prev.Next, tail = prev.Next, tail, tail.Next
			}
		}
		prev = tail
	}
	return hd.Next
}
