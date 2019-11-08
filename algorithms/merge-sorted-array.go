package algorithms

// https://leetcode.com/problems/merge-sorted-array/description/

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	for i >= 0 && n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}
}

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

