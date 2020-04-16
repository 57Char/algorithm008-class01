package remove_duplicates_from_sorted_array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for _, val := range nums {
		if nums[i] != val {
			i++
			nums[i] = val
		}
	}
	return i + 1
}

func removeDuplicatesV2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	i, j := 1, 1
	for ; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
