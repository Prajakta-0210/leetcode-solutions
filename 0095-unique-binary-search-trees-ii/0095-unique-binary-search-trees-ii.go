/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return build(1, n)
}

func build(start, end int) []*TreeNode {
	var result []*TreeNode

	// No nodes in this range
	if start > end {
		return []*TreeNode{nil}
	}

	// Try every value as the root
	for root := start; root <= end; root++ {

		// Generate all possible left subtrees
		leftTrees := build(start, root-1)

		// Generate all possible right subtrees
		rightTrees := build(root+1, end)

		// Combine every left subtree with every right subtree
		for _, left := range leftTrees {
			for _, right := range rightTrees {

				rootNode := &TreeNode{
					Val:   root,
					Left:  left,
					Right: right,
				}

				result = append(result, rootNode)
			}
		}
	}

	return result
}