package algorithms

// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n - 1
	for left <= right {
		mid := left + ((right-left) >> 1)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
