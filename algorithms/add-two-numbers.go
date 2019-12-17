package algorithms

// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	curr := prev
	sum := 0

	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		sum /= 10
	}
	return prev.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	curr := prev
	carry := 0

	for l1 != nil || l2 != nil{
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry == 1 {
		curr.Next = &ListNode{Val: carry}
	}

	return prev.Next
}
