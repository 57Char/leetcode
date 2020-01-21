package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		Name   string
		Array  []int
		Expect []int
	}{
		{
			Name:   "0 element",
			Array:  []int{},
			Expect: []int{},
		},
		{
			Name:   "1 element",
			Array:  []int{1},
			Expect: []int{1},
		},
		{
			Name:   "2 elements",
			Array:  []int{2, 1},
			Expect: []int{1, 2},
		},
		{
			Name:   "3 elements",
			Array:  []int{3, 1, 3},
			Expect: []int{1, 3, 3},
		},
		{
			Name:   "N elements",
			Array:  []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			Expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			QuickSort(test.Array, len(test.Array))
			assert.Equal(t, test.Expect, test.Array)
		})
	}
}
