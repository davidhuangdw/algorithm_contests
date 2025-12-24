package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// clone & couple
	cur := head
	for cur != nil {
		clone := &Node{cur.Val, cur.Next, cur.Random}
		cur.Next = clone
		cur = clone.Next
	}

	// update clone's 'random'
	for cur := head; cur != nil; cur = cur.Next.Next {
		clone := cur.Next
		if clone.Random != nil {
			clone.Random = clone.Random.Next
		}
	}

	// decouple
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
