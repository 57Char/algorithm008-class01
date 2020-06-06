package intersection_of_two_arrays_ii

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/

// 执行用时 : 4 ms, 在所有 Go 提交中击败了 90.87% 的用户
// 内存消耗 : 2.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}
	i, j := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	intersect := []int{}
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			intersect = append(intersect, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return intersect
}

// 执行用时 : 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 3.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
func intersectV2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersectV2(nums2, nums1)
	}
	m := map[int]int{}
	for _, n := range nums1 {
		m[n] += 1
	}
	intersect := []int{}
	for _, n := range nums2 {
		if m[n] > 0 {
			intersect = append(intersect, n)
			m[n] -= 1
		}
	}
	return intersect
}