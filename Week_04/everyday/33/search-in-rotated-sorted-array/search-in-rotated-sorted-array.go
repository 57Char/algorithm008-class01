package search_in_rotated_sorted_array

// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		if target == nums[l] {
			return l
		}
		if target == nums[r] {
			return r
		}
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		}
		if nums[l] > nums[r] {
			l++
			r--
		} else {
			if target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
