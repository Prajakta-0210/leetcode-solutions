/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    // Store inorder value -> index
    inorderMap := make(map[int]int)

    for i, val := range inorder {
        inorderMap[val] = i
    }

    preorderIndex := 0

    var build func(left, right int) *TreeNode

    build = func(left, right int) *TreeNode {
        // No elements in this subtree
        if left > right {
            return nil
        }

        // First element in preorder is the root
        rootValue := preorder[preorderIndex]
        preorderIndex++

        root := &TreeNode{Val: rootValue}

        // Find root position in inorder
        mid := inorderMap[rootValue]

        // Build left subtree
        root.Left = build(left, mid-1)

        // Build right subtree
        root.Right = build(mid+1, right)

        return root
    }

    return build(0, len(inorder)-1)
}