package algorithms

// https://leetcode.com/problems/rotate-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	copyHead := head
	size := 1
	for copyHead.Next != nil {
		copyHead = copyHead.Next
		size++
	}

	// make a cycle
	copyHead.Next = head
	// 5->1->2-3->4->5->1->2->3->4...

	for i := size - k % size; i > 1; i-- {
		head = head.Next
	}
	copyHead = head.Next
	// 4->5->1->2->3->4...
	head.Next = nil
	// 4->5->1->2->3

	return copyHead
}

/*func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	count := 1
	for curr.Next != nil {
		curr = curr.Next
		count++
	}
	curr.Next = head
	// TODO K=0时有问题
	for i:=0; i<k%count; i++ {
		head = head.Next
	}
	curr = head.Next
	head.Next = nil
	return curr
}*/