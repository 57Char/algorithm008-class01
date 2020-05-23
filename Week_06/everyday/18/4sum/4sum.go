package _sum

import "sort"

// https://leetcode.com/problems/4sum/

// Runtime: 8 ms, faster than 91.91% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 4Sum.
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}

// Runtime: 4 ms, faster than 98.53% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 4Sum
func fourSumV2(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// repeat
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// min
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// max
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// repeat
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// min
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// max
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}

// Runtime: 4 ms, faster than 98.53% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 4Sum.
func fourSumV3(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// repeat
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// min
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// max
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// repeat
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// min
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// max
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				} else {
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				}
			}
		}
	}
	return res
}
