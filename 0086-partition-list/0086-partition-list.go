func partition(head *ListNode, x int) *ListNode {
	
	lessDummy := &ListNode{}
	greaterDummy := &ListNode{}

	less := lessDummy
	greater := greaterDummy

	current := head

	for current != nil {
		if current.Val < x {
			less.Next = current
			less = less.Next
		} else{
			greater.Next = current
			greater = greater.Next
		}

		current = current.Next
	}

	greater.Next = nil

	less.Next = greaterDummy.Next

	return lessDummy.Next
}