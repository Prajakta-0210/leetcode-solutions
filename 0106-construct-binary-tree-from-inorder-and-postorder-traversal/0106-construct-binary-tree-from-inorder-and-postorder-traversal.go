/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {

    // Store inorder value -> index
    indexMap := make(map[int]int)

    for i, val := range inorder {
        indexMap[val] = i
    }

    // postIndex points to the current root in postorder
    postIndex := len(postorder) - 1

    var build func(left, right int) *TreeNode

    build = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }

        // Last element of postorder is the root
        rootValue := postorder[postIndex]
        postIndex--

        root := &TreeNode{Val: rootValue}

        rootIndex := indexMap[rootValue]

        // IMPORTANT:
        // Since we are moving backwards in postorder,
        // we must build RIGHT first, then LEFT.
        root.Right = build(rootIndex+1, right)
        root.Left = build(left, rootIndex-1)

        return root
    }

    return build(0, len(inorder)-1)
}