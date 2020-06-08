package permutations

// https://leetcode.com/problems/permutations/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, []int{}, &res)
	return res
}

func backtrack(nums []int, track []int, res *[][]int) {
	if len(nums) == len(track) {
		*res = append(*res, append([]int{}, track...))
		return
	}
	for _, n := range nums {
		if contains(track, n) {
			continue
		}
		backtrack(nums, append(track, n), res)
	}
}

func contains(track []int, n int) bool {
	for _, v := range track {
		if v == n {
			return true
		}
	}
	return false
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
// Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Permutations.
func permuteV2(nums []int) [][]int {
	res := [][]int{}
	backtrackV2(nums, 0, len(nums), &res)
	return res
}

func backtrackV2(nums []int, level, n int, res *[][]int) {
	if level == n-1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := level; i < n; i++ {
		nums[i], nums[level] = nums[level], nums[i]
		backtrackV2(nums, level+1, n, res)
		nums[i], nums[level] = nums[level], nums[i]
	}
}