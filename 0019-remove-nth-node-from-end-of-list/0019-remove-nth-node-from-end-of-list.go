func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Dummy node handles edge cases (e.g., removing the head)
	dummy := &ListNode{Next: head}

	fast := dummy
	slow := dummy

	// Move fast pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node from the end
	slow.Next = slow.Next.Next

	return dummy.Next
}