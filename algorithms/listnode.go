package algorithms

import (
	"fmt"
)

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
