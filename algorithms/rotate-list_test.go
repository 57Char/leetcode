package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	head := NewListNode(1, 2, 3, 4, 5)
	expect := NewListNode(4, 5,1,2,3).String()
	assert.Equal(t, expect, rotateRight(head, 2).String())
}

