package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		T      string
		Expect bool
	}{
		{
			"S is empty and T is not empty",
			"",
			"abc",
			false,
		},
		{
			"S is not empty and T is empty",
			"abc",
			"",
			false,
		},
		{
			"len(S) != len(T)",
			"abc",
			"abcd",
			false,
		},
		{
			"len(S) == len(T), but is not anagram",
			"abcd",
			"abce",
			false,
		},
		{
			"len(S) == len(T), but is anagram",
			"abcd",
			"dcba",
			true,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, isAnagram(test.S, test.T))
		})
	}
}

func TestIsAnagramV2(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		T      string
		Expect bool
	}{
		{
			"S is empty and T is not empty",
			"",
			"abc",
			false,
		},
		{
			"S is not empty and T is empty",
			"abc",
			"",
			false,
		},
		{
			"len(S) != len(T)",
			"abc",
			"abcd",
			false,
		},
		{
			"len(S) == len(T), but is not anagram",
			"abcd",
			"abce",
			false,
		},
		{
			"len(S) == len(T), but is anagram",
			"abcd",
			"dcba",
			true,
		},
		{
			"Unicode",
			"你好世界",
			"界世好你",
			true,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, isAnagramV2(test.S, test.T))
		})
	}
}
