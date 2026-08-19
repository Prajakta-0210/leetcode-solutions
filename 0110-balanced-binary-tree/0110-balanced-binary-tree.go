/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return height(root) != -1
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftHeight := height(root.Left)
    if leftHeight == -1 {
        return -1
    }

    rightHeight := height(root.Right)
    if rightHeight == -1 {
        return -1
    }

    // Difference between left and right subtree heights > 1
    if abs(leftHeight-rightHeight) > 1 {
        return -1
    }

    // Return height of current subtree
    if leftHeight > rightHeight {
        return leftHeight + 1
    }

    return rightHeight + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}