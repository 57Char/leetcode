package algorithms

// https://leetcode.com/problems/rotate-array/description/

func rotate(nums []int, k int) {
	l := len(nums)
	if l <= 1 || k <= 0 {
		return
	}

	k = k % l
	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		for j := l - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

func rotateV2(nums []int, k int) {
	l := len(nums)
	if l <= 1 || k <= 0 {
		return
	}

	k = k % l
	if k == 0 {
		return
	}

	newNums := make([]int, l)
	copy(newNums[:k], nums[l-k:])
	copy(newNums[k:], nums[:l-k])
	copy(nums, newNums)
}

func rotateV3(nums []int, k int) {
	l := len(nums)
	if l <= 1 || k <= 0 {
		return
	}

	k = k % l
	if k == 0 {
		return
	}

	copy(nums, append(nums[l-k:], nums[:l-k]...))
}