package algorithms

// https://leetcode.com/problems/validate-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node, lower, upper *TreeNode) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if lower != nil && val <= lower.Val {
		return false
	}
	if upper != nil && val >= upper.Val {
		return false
	}
	if !helper(node.Left, lower, node) {
		return false
	}
	if !helper(node.Right, node, upper) {
		return false
	}
	return true
}

func isValidBSTV2(root *TreeNode) bool {
	h := &Helper{}
	return h.Help(root)
}

type Helper struct {
	Pre *TreeNode
}

func (h *Helper) Help(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !h.Help(root.Left) {
		return false
	}
	if h.Pre != nil && h.Pre.Val >= root.Val {
		return false
	}
	h.Pre = root
	return h.Help(root.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBSTV3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*Group{&Group{root, nil, nil}}
	for len(stack) > 0 {
		group := stack[0]
		stack = stack[1:]
		node, lower, upper := group.Node, group.Lower, group.Upper
		if node == nil {
			continue
		}
		val := node.Val
		if lower != nil && val <= lower.Val {
			return false
		}
		if upper != nil && val >= upper.Val {
			return false
		}
		stack = append(stack, &Group{node.Left, lower, node})
		stack = append(stack, &Group{node.Right, node, upper})
	}

	return true
}

type Group struct {
	Node, Lower, Upper *TreeNode
}