package algorithms

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	level := []*TreeNode{root}
	for len(level) > 0 {
		depth++
		q := []*TreeNode{}
		for _, node := range level {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level = q
	}
	return depth
}

