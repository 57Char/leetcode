package algorithms

import "fmt"

func permutation(nums []int, start int) {
	if start == len(nums) - 1 {
		fmt.Printf("%v\n", nums)
	}
	for i:=start; i<len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permutation(nums, start+1)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
