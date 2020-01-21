package algorithms

import (
	"testing"
)

func TestPermutation(t *testing.T)  {
	permutation([]int{1,2,3,4},0)
}

//[1 2 3 4]
//[1 2 4 3]
//[1 3 2 4]
//[1 3 4 2]
//[1 4 3 2]
//[1 4 2 3]
//[2 1 3 4]
//[2 1 4 3]
//[2 3 1 4]
//[2 3 4 1]
//[2 4 3 1]
//[2 4 1 3]
//[3 2 1 4]
//[3 2 4 1]
//[3 1 2 4]
//[3 1 4 2]
//[3 4 1 2]
//[3 4 2 1]
//[4 2 3 1]
//[4 2 1 3]
//[4 3 2 1]
//[4 3 1 2]
//[4 1 3 2]
//[4 1 2 3]