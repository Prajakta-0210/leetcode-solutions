/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) [][]int {
    result := [][]int{}
    path := []int{}

    dfs(root, targetSum, path, &result)

    return result
}

func dfs(root *TreeNode, targetSum int, path []int, result *[][]int) {
    if root == nil {
        return
    }

    // Add current node to path
    path = append(path, root.Val)

    // Check if current node is a leaf
    if root.Left == nil && root.Right == nil {
        if root.Val == targetSum {
            // Make a copy of path
            currentPath := append([]int{}, path...)
            *result = append(*result, currentPath)
        }
        return
    }

    // Remaining sum after taking current node
    remaining := targetSum - root.Val

    // Explore left and right subtrees
    dfs(root.Left, remaining, path, result)
    dfs(root.Right, remaining, path, result)
}