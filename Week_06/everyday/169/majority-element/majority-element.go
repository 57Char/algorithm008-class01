package majority_element

import "sort"

// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementV2(nums []int) int {
	count, majority := 0, 0

	for _, n := range nums {
		if count == 0 {
			majority = n
		}
		if n == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}