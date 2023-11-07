package tree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return arrayEquals(getLeaves(root1), getLeaves(root2))
}

func getLeaves(root *TreeNode) []int {
	var stack []*TreeNode
	var leaves []int
	stack = append(stack, root)

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Left == nil && top.Right == nil {
			leaves = append(leaves, top.Val)
		}
	}
	return leaves
}

func arrayEquals(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
