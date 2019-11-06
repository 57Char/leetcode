package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		Name   string
		K      int
		Input  []int
		Expect []int
	}{
		{
			"Empty Input",
			10,
			[]int{},
			[]int{},
		},
		{
			"Empty K",
			0,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) == K",
			3,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) * 2 == K",
			6,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) > K",
			3,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"Len(Input) < K",
			8,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotate(test.Input, test.K)
			assert.Equal(t, test.Expect, test.Input)
		})
	}
}

func TestRotateV2(t *testing.T) {
	cases := []struct {
		Name   string
		K      int
		Input  []int
		Expect []int
	}{
		{
			"Empty Input",
			10,
			[]int{},
			[]int{},
		},
		{
			"Empty K",
			0,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) == K",
			3,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) * 2 == K",
			6,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) > K",
			3,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"Len(Input) < K",
			8,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotateV2(test.Input, test.K)
			assert.Equal(t, test.Expect, test.Input)
		})
	}
}

func TestRotateV3(t *testing.T) {
	cases := []struct {
		Name   string
		K      int
		Input  []int
		Expect []int
	}{
		{
			"Empty Input",
			10,
			[]int{},
			[]int{},
		},
		{
			"Empty K",
			0,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) == K",
			3,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) * 2 == K",
			6,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			"Len(Input) > K",
			3,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"Len(Input) < K",
			8,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotateV3(test.Input, test.K)
			assert.Equal(t, test.Expect, test.Input)
		})
	}
}