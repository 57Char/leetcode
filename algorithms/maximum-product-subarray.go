package algorithms

import "math"

// https://leetcode.com/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, maxVal, minVal := math.MinInt64, 1, 1
	for _, n := range nums {
		if n < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(maxVal*n, n)
		minVal = min(minVal*n, n)
		res = max(res, maxVal)
	}
	return res
}

func maxProductV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, min, res := nums[0], nums[0], nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}

		if v < 0 {
			max, min = min, max
		}

		if max * v > v {
			max = max * v
		} else {
			max = v
		}

		if min * v < v {
			min = min * v
		} else {
			min = v
		}

		if max > res {
			res = max
		}
	}

	return res
}