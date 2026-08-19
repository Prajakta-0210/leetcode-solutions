/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
    current := root

    for current != nil {
        // If there is no left subtree,
        // move to the right subtree.
        if current.Left == nil {
            current = current.Right
            continue
        }

        // Find the rightmost node of the left subtree.
        predecessor := current.Left

        for predecessor.Right != nil {
            predecessor = predecessor.Right
        }

        // Connect the original right subtree
        // to the rightmost node of the left subtree.
        predecessor.Right = current.Right

        // Move the left subtree to the right.
        current.Right = current.Left

        // Remove the left pointer.
        current.Left = nil

        // Move to the next node.
        current = current.Right
    }
}