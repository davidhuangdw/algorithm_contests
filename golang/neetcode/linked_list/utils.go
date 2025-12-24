package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{}
	current := head
	for _, v := range values {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head.Next
}

func listToSlice(head *ListNode) []int {
	result := make([]int, 0)
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
