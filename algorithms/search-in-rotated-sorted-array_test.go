package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Target int
		Expect int
	}{
		{
			"In List",
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			"In List",
			[]int{4, 5, 6, 7, 0, 1, 2},
			4,
			0,
		},
		{
			"In List",
			[]int{4, 5, 6, 7, 0, 1, 2},
			7,
			3,
		},
		{
			"In List",
			[]int{4, 5, 6, 7, 0, 1, 2},
			2,
			6,
		},
		{
			"Not In List",
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			"Not In List",
			[]int{},
			3,
			-1,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, search(test.Nums, test.Target))
		})
	}
}
