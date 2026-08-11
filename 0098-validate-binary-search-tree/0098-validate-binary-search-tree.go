/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}

	// Current node must be strictly inside the valid range
	if int64(root.Val) <= min || int64(root.Val) >= max {
		return false
	}

	// Left subtree: values must be smaller
	// Right subtree: values must be greater
	return validate(root.Left, min, int64(root.Val)) &&
		validate(root.Right, int64(root.Val), max)
}