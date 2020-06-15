package reverse_pairs

// https://leetcode.com/problems/reverse-pairs/

// TODO 需要优化
func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	return mergeSort(nums, 0, n-1)
}

func mergeSort(nums []int, left, right int) int {
	if right <= left {
		return 0
	}
	mid := left + (right-left)>>1 // (left + right) / 2
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	count += merge(nums, left, mid, right)
	return count
}

func merge(nums []int, left, mid, right int) int {
	count := 0
	var merge []int
	i, j, p := left, mid+1, mid+1

	for i <= mid {
		for p <= right && nums[i] > nums[p]*2 {
			p++
		}
		count += p - (mid + 1)
		for j <= right && nums[i] >= nums[j] {
			merge = append(merge, nums[j])
			j++
		}
		merge = append(merge, nums[i])
		i++
	}
	for j <= right {
		merge = append(merge, nums[j])
		j++
	}

	_ = append(nums[:left], merge...)
	return count
}
