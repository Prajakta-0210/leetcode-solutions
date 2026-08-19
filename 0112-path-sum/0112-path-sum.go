/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    // If this is a leaf, check whether its value
    // matches the remaining target
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    // Subtract current node's value
    remaining := targetSum - root.Val

    return hasPathSum(root.Left, remaining) ||
        hasPathSum(root.Right, remaining)
}