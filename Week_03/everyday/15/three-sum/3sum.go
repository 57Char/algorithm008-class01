package three_sum

import "sort"

// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i, v := range nums {
		if v > 0 {
			return res
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// very slow
func threeSumV2(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i, x := range nums[:n-2] {
		if x > 0 {
			return res
		}
		if i > 0 && x == nums[i-1] {
			continue
		}
		m := map[int]bool{}
		for j := i + 1; j < n; j++ {
			y := nums[j]
			if m[y] {
				res = append(res, []int{x, -x - y, y})
				for j+1 < n && y == nums[j+1] {
					j++
				}
			} else {
				m[-x-y] = true
			}
		}
	}
	return res
}

