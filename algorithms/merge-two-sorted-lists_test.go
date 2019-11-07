package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		Name   string
		L1     *ListNode
		L2     *ListNode
		Expect string
	}{
		{
			"Empty L1",
			nil,
			&ListNode{},
			(&ListNode{}).String(),
		},
		{
			"Empty L2",
			&ListNode{},
			nil,
			(&ListNode{}).String(),
		},
		{
			"Empty L1 and Empty L2",
			nil,
			nil,
			"",
		},
		{
			"Merge L1 and L2 case one",
			NewListNode(1, 2),
			NewListNode(1, 2),
			NewListNode(1, 1, 2, 2).String(),
		},
		{
			"Merge L1 and L2 case two",
			NewListNode(1, 2, 3, 4),
			NewListNode(2, 3, 5, 6),
			NewListNode(1, 2, 2, 3, 3, 4, 5, 6).String(),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			merged := mergeTwoLists(test.L1, test.L2)
			if merged != nil {
				assert.Equal(t, test.Expect, merged.String())
			}
		})
	}
}
