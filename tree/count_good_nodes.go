package tree

func goodNodes(root *TreeNode) int {
	return goodNodesRecursion(root, -10001)
}

func goodNodesRecursion(node *TreeNode, maxOfPath int) int {
	if node == nil {
		return 0
	}
	if node.Val >= maxOfPath {
		return goodNodesRecursion(node.Left, node.Val) + 1 + goodNodesRecursion(node.Right, node.Val)
	}
	return goodNodesRecursion(node.Left, maxOfPath) + goodNodesRecursion(node.Right, maxOfPath)
}
