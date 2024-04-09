package tree

var max = 0

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	zigzag(root, true, -1)
	zigzag(root, false, -1)
	return max
}

func zigzag(node *TreeNode, fromRight bool, count int) {
	if node == nil {
		if count > max {
			max = count
		}
		return
	}
	if fromRight {
		zigzag(node.Left, false, count+1)
		zigzag(node.Right, false, -1)
		return
	}
	zigzag(node.Right, true, count+1)
	zigzag(node.Left, true, -1)
}
