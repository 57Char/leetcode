package algorithms

// https://leetcode.com/problems/reverse-linked-list/
// https://leetcode.com/problems/reverse-linked-list/solution/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// Time complexity: O(n)
// Space complexity: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// Time complexity: O(n)
// Space complexity: O(n)


