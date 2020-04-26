package sort_colors

// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j, k := 0, 0, n-1
	for i <= k {
		switch nums[i] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			i++
		case 2:
			nums[i], nums[k] = nums[k], nums[i]
			k--
		}
	}
}
