/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    // Start from the first node of each level
    leftmost := root

    for leftmost.Left != nil {
        current := leftmost

        for current != nil {
            // Connect left child to right child
            current.Left.Next = current.Right

            // Connect right child to next node's left child
            if current.Next != nil {
                current.Right.Next = current.Next.Left
            }

            // Move to next node in the current level
            current = current.Next
        }

        // Move to the next level
        leftmost = leftmost.Left
    }

    return root
}