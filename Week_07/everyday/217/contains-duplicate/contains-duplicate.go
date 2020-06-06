package contains_duplicate

import "sort"

// https://leetcode.com/problems/contains-duplicate/

// 执行用时 : 24 ms, 在所有 Go 提交中击败了 83.67% 的用户
// 内存消耗 : 6.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
func containsDuplicate(nums []int) bool {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 执行用时 : 20 ms, 在所有 Go 提交中击败了 97.11% 的用户
// 内存消耗 : 6.3 MB, 在所有 Go 提交中击败了 100.00% 的用户
func containsDuplicateV2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	visited := map[int]bool{}
	for _, n := range nums {
		if visited[n] {
			return true
		}
		visited[n] = true
	}
	return false
}
