/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{Next: head}
	prevGroup := dummy

	for {

		// Check if k nodes exist
		kth := prevGroup
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}

		if kth == nil {
			break
		}

		groupNext := kth.Next

		// Reverse group
		prev := groupNext
		curr := prevGroup.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// Connect
		temp := prevGroup.Next
		prevGroup.Next = kth
		prevGroup = temp
	}

	return dummy.Next
}