package algorithms

// https://leetcode.com/problems/merge-two-sorted-lists/description/

import (
	"fmt"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func NewListNode(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	v := vals[0]

	l := &ListNode{
		Val: v,
	}

	if len(vals) > 1 {
		l.Next = NewListNode(vals[1:]...)
	}

	return l
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	vals := []int{l.Val}

	ln := l
	for {
		ln = ln.Next
		if ln == nil {
			break
		}
		vals = append(vals, ln.Val)
	}

	return fmt.Sprintf("%v", vals)
}
