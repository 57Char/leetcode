package algorithms

// https://leetcode.com/problems/invert-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

func invertTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	level := []*TreeNode{root}
	for len(level) > 0 {
		q := []*TreeNode{}
		for _, node := range level {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
			//fmt.Printf("Inverted:%v->%v->%v\n",node.Val, node.Left, node.Right)
		}
		//for _, node := range q {
		//	fmt.Printf("%v->%v->%v\n",node.Val, node.Left, node.Right)
		//}
		level = q
	}
	return root
}