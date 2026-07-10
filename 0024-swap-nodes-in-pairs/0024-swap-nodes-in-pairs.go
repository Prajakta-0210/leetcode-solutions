/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {

    dummy := &ListNode{}
    dummy.Next = head

    current := dummy

    for current.Next != nil && current.Next.Next != nil {

        first := current.Next
        second := current.Next.Next

        // Swap
        first.Next = second.Next
        second.Next = first
        current.Next = second

        // Move to next pair
        current = first
    }

    return dummy.Next
}