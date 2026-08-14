/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
    // Both nodes are empty
    if left == nil && right == nil {
        return true
    }

    // Only one node is empty
    if left == nil || right == nil {
        return false
    }

    // Values are different
    if left.Val != right.Val {
        return false
    }

    // Compare opposite sides
    return isMirror(left.Left, right.Right) &&
           isMirror(left.Right, right.Left)
}