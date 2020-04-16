package rotate_array

// https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	n := len(nums)
	if n <= 1 || k <= 0  {
		return
	}
	k %= n
	if k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotateV2(nums []int, k int) {
	n := len(nums)
	if n <= 1 || k <= 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	newNums := make([]int, n)
	copy(newNums[:k], nums[n-k:])
	copy(newNums[k:], nums[:n-k])
	copy(nums, newNums)
}

func rotateV3(nums []int, k int) {
	n := len(nums)
	if n <= 1 || k <= 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	copy(nums, append(nums[n-k:], nums[:n-k]...))
}

func rotateV4(nums []int, k int) {
	n := len(nums)
	if n <= 1 || k <= 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
