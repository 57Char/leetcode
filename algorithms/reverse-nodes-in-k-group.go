package algorithms

// https://leetcode.com/problems/reverse-nodes-in-k-group/submissions/
// https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/11423/Short-but-recursive-Java-code-with-comments
// https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/11440/Non-recursive-Java-solution-and-idea

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for curr != nil && count < k {
		curr = curr.Next
		count ++
	}
	if count < k {
		return head
	}
	curr = reverseKGroup(curr, k)
	count = 0
	for count < k {
		next := head.Next
		head.Next = curr
		curr = head
		head = next
		count ++
	}
	head = curr
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// TODO 这段代码有问题：[1,2]2, [1,2,3,4,5]2
//func reverseKGroupV2(head *ListNode, k int) *ListNode {
//	curr := head
//	i := 0
//	for i < k && curr != nil {
//		curr = curr.Next
//		i++
//	}
//	if i < k {
//		//fmt.Printf("i:%d\n",i)
//		return head
//	}
//	if curr == nil {
//		return reverse(head)
//	}
//	next := curr.Next
//	curr.Next = nil
//
//	newHead := reverse(head)
//	newNext := reverseKGroupV2(next, k)
//	//fmt.Printf("newHead: %v, newNext: %v",newHead, newNext)
//	head.Next = newNext
//
//	return newHead
//}
//
//func reverse(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	p := reverse(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return p
//}