package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node { // insert to 'Next'
	if head == nil {
		return nil
	}

	// insert clone into 'Next': each origin know its clone
	cur := head
	for cur != nil {
		clone := &Node{cur.Val, cur.Next, cur.Random}
		cur.Next = clone
		cur = clone.Next
	}

	// fix clone's 'random': origin unchanged, clone's random == origin-random's clone
	for cur := head; cur != nil; cur = cur.Next.Next {
		clone := cur.Next
		if clone.Random != nil {
			clone.Random = clone.Random.Next
		}
	}

	// fix clone's 'next' & origin: next-origin unfixed-yet, clone's next == next-origin's clone
	hd := head.Next
	l, r := head, hd
	for cur = hd.Next; cur != nil; {
		l.Next = cur
		l = cur
		cur = cur.Next

		r.Next = cur
		r = cur
		cur = cur.Next
	}
	l.Next = nil

	return hd
}

func copyRandomList(head *Node) *Node { // insert to 'Random'
	if head == nil {
		return nil
	}

	// insert clone into 'Random': each origin know its clone
	for cur := head; cur != nil; cur = cur.Next {
		cur.Random = &Node{cur.Val, cur.Random, cur.Random}
	}

	// fix clone's 'random': origin unchanged, clone's random == origin-random's clone
	for cur := head; cur != nil; cur = cur.Next {
		clone := cur.Random
		if clone.Random != nil {
			clone.Random = clone.Random.Random
		}
	}

	// fix clone's 'next' & origin: next-origin unfixed-yet, clone's next == next-origin's clone
	hd := head.Random
	for cur := head; cur != nil; cur = cur.Next {
		clone := cur.Random
		cur.Random = clone.Next
		if cur.Next != nil {
			clone.Next = cur.Next.Random
		} else {
			clone.Next = nil
		}
	}
	return hd
}
