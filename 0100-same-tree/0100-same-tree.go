/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both are empty
	if p == nil && q == nil {
		return true
	}

	// One is empty, the other is not
	if p == nil || q == nil {
		return false
	}

	// Values are different
	if p.Val != q.Val {
		return false
	}

	// Compare left and right subtrees
	return isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}