package intersection_of_two_arrays_ii

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/

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

func intersectV2(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	m := map[int]int{}
	intersect := make([]int, 0, n1)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if v2, ok := m[v]; ok && v2 > 0 {
			intersect = append(intersect, v)
			m[v]--
		}
	}
	return intersect
}

func intersectV3(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	m := map[int]int{}
	intersect := make([]int, 0, n1)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v]--; m[v] >= 0 {
			intersect = append(intersect, v)
		}
	}
	return intersect
}
