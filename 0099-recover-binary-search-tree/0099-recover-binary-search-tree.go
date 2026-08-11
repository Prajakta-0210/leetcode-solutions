/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Traverse left subtree
		inorder(node.Left)

		// Check if inorder order is broken
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}

		prev = node

		// Traverse right subtree
		inorder(node.Right)
	}

	inorder(root)

	// Swap the two incorrect values
	first.Val, second.Val = second.Val, first.Val
}