package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"

)

func TestBinarySearch(t *testing.T) {
	a := []int{8, 11, 19, 23, 27, 33, 45, 55, 67, 98}
	n := len(a)
	assert.Equal(t, 2, binarySearch(a, n, 19))
	assert.Equal(t, 9, binarySearch(a, n, 98))
	assert.Equal(t, -1, binarySearch(a, n, 199))
}
