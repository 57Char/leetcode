package algorithms

// https://leetcode.com/problems/path-sum-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return append(res, []int{root.Val})
		}
		return res
	}
	for _, path := range pathSum(root.Left, sum - root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	for _, path := range pathSum(root.Right, sum - root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	return res
}
