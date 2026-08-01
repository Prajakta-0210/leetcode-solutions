func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	current := prev.Next

	for i := 0; i < right-left; i++ {
		nextNode := current.Next

		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}