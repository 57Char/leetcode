package algorithms

// https://leetcode.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		p1, p2 = p1.Next, p2.Next
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
	}
	return p1
}

func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
