package move_zeroes

// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i, val := range nums {
		if val != 0 {
			nums[i], nums[j] = 0, val
			j++
		}
	}
}

func moveZeroesV2(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i, val := range nums {
		if val != 0 {
			if i != j {
				nums[i] = 0
			}
			nums[j] = val
			j++
		}
	}
}