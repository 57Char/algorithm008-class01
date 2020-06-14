package merge_sorted_array

import (
	"sort"
)

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for i >= 0 && n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}
}

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for m >= 0 && n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}
}

func mergeV3(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}
	for n > 0 {
		nums1[i] = nums2[n-1]
		n--
		i--
	}
}

func mergeV4(nums1 []int, m int, nums2 []int, n int) {
	//more := len(nums1) - m - n
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	//if more > 0 {
	//	nums3 := make([]int, more)
	//	nums1 = append(nums1, nums3...)
	//}
}