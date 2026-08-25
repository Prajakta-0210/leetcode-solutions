func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	curr := root

	for curr != nil {
		var nextLevelStart *Node
		var prev *Node

		// Traverse current level
		for curr != nil {

			if curr.Left != nil {
				if nextLevelStart == nil {
					nextLevelStart = curr.Left
				}

				if prev != nil {
					prev.Next = curr.Left
				}

				prev = curr.Left
			}

			if curr.Right != nil {
				if nextLevelStart == nil {
					nextLevelStart = curr.Right
				}

				if prev != nil {
					prev.Next = curr.Right
				}

				prev = curr.Right
			}

			curr = curr.Next
		}

		// Move to next level
		curr = nextLevelStart
	}

	return root
}