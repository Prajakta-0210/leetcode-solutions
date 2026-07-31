func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			duplicateValue := current.Val

			for current != nil && current.Val == duplicateValue {
				current = current.Next
			}

			prev.Next = current
		} else {
			prev = current
			current = current.Next
		}
	}

	return dummy.Next
}