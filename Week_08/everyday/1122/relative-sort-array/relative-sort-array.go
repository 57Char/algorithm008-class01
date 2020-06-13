package relative_sort_array

import "sort"

// https://leetcode.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	i, n := 0, len(arr1)
	for _, v := range arr2 {
		for j := 0; j < n; j++ {
			if v == arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
				i++
			}
		}
	}
	sort.Ints(arr1[i:])
	//return append(arr1[:i], arr1[i:]...)
	return arr1
}