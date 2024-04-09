package tree

import "math"

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := []*TreeNode{root}
	maxSum := math.MinInt
	ans := -1
	curLevel := 0
	for len(level) > 0 {
		curLevel++
		curLevelSum := 0
		var nextLevel []*TreeNode
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
			curLevelSum += node.Val
		}
		if curLevelSum > maxSum {
			maxSum = curLevelSum
			ans = curLevel
		}
		level = nextLevel
	}
	return ans
}
