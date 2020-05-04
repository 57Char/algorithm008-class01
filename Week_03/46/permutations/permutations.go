package permutations

// https://leetcode.com/problems/permutations/

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
