package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	head := NewListNode(1, 2, 3, 4, 5)
	node := head.Next.Next // delete node 3
	expect := NewListNode(1, 2, 4, 5).String()
	deleteNode(node)
	assert.Equal(t, expect, head.String())
}
