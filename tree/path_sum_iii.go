package tree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return recursion(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func recursion(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	targetSum -= root.Val
	if targetSum == 0 {
		return 1 + recursion(root.Left, targetSum) + recursion(root.Right, targetSum)

	}
	return recursion(root.Left, targetSum) + recursion(root.Right, targetSum)
}
