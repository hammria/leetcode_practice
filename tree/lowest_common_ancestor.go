package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	foundLeft := lowestCommonAncestor(root.Left, p, q)
	foundRight := lowestCommonAncestor(root.Right, p, q)
	if foundLeft != nil && foundRight != nil {
		return root
	}
	if foundLeft != nil {
		return foundLeft
	}
	return foundRight
}
