package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var level []*TreeNode
	level = append(level, root)
	ans = append(ans, root.Val)
	for {
		var nextLevel []*TreeNode
		for i := range level {
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		if len(nextLevel) > 0 {
			ans = append(ans, nextLevel[len(nextLevel)-1].Val)
		} else {
			break
		}
		level = nextLevel
	}
	return ans
}
