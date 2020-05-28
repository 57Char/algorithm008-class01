package combinations

// https://leetcode.com/problems/combinations/

func combine(n int, k int) [][]int {
	// TODO n, k 的边界条件处理
	res := [][]int{}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	backtrack(nums, []int{}, k, &res)
	return res
}

func backtrack(nums, track []int, k int, res *[][]int) {
	if len(track) == k {
		*res = append(*res, append([]int{}, track...))
		return
	}
	if len(nums)+len(track) < k {
		return
	}
	for i, v := range nums {
		backtrack(nums[i+1:], append(track, v), k, res)
	}
}

func combineV2(n int, k int) [][]int {
	res := [][]int{}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	backtrackV2(nums, []int{}, 0, n, k, &res)
	return res
}

func backtrackV2(nums, track []int, i, n, k int, res *[][]int) {
	if len(track) == k {
		*res = append(*res, append([]int{}, track...))
		return
	}
	if n-i+len(track) < k {
		return
	}
	for ; i < n; i++ {
		backtrackV2(nums, append(track, nums[i]), i+1, n, k, res)
	}
}