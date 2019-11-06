package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates([]int{}))
	assert.Equal(t, 1, removeDuplicates([]int{1}))
	assert.Equal(t, 1, removeDuplicates([]int{1, 1}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 2}))
	assert.Equal(t, 3, removeDuplicates([]int{1, 1, 2, 2, 3, 3, 3}))
	assert.Equal(t, 4, removeDuplicates([]int{1, 2, 3, 4, 4, 4, 4, 4}))
}
