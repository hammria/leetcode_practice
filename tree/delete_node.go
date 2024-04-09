package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			root.Val = findMinNode(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findMinNode(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}
