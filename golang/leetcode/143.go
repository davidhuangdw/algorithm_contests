package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// end of half
	s, f := head, head.Next
	for f.Next != nil && f.Next.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	l, r := head, s.Next
	s.Next = nil

	// reverse r
	p := r
	r = nil
	for p != nil {
		p.Next, r, p = r, p, p.Next
	}

	// inter-merge l, r
	hd := &ListNode{}
	p = hd
	for l != nil || r != nil {
		if l != nil {
			p.Next, l, p = l, l.Next, l
		}
		if r != nil {
			p.Next, r, p = r, r.Next, r
		}
	}
	p.Next = nil
}
