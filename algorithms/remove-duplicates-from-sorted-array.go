package algorithms

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slowP := 0
	for _, val := range nums {
		if nums[slowP] != val {
			slowP++
			nums[slowP] = val
		}
	}
	return slowP + 1
}
