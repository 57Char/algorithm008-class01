package permutations_ii

import "sort"

// https://leetcode.com/problems/permutations-ii/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations II.
// Memory Usage: 3.5 MB, less than 100.00% of Go online submissions for Permutations II.
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, len(nums), &res)
	return res
}

func backtrack(nums []int, level, n int, res *[][]int) {
	if level == n-1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	visited := map[int]bool{}
	for i := level; i < n; i++ {
		if visited[nums[i]] {
			continue
		}
		nums[i], nums[level] = nums[level], nums[i]
		backtrack(nums, level+1, n, res)
		nums[i], nums[level] = nums[level], nums[i]
		visited[nums[i]] = true
	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations II.
// Memory Usage: 3.5 MB, less than 100.00% of Go online submissions for Permutations II.
func permuteUniqueV2(nums []int) [][]int {
	res, n := [][]int{}, len(nums)
	visited := make([]bool, n)
	sort.Ints(nums)
	backtrackV2(nums, []int{}, visited, n, &res)
	return res
}

func backtrackV2(nums, track []int, visited []bool, n int, res *[][]int) {
	if len(track) == n {
		*res = append(*res, append([]int{}, track...))
		return
	}
	for i, v := range nums {
		if visited[i] || (i > 0 && v == nums[i-1] && !visited[i-1]) {
			continue
		}
		visited[i] = true
		backtrackV2(nums, append(track, v), visited, n, res)
		visited[i] = false
	}
}

// Runtime: 4 ms, faster than 65.54% of Go online submissions for Permutations II.
// Memory Usage: 3.6 MB, less than 100.00% of Go online submissions for Permutations II.
func permuteUniqueV3(nums []int) [][]int {
	res, n := [][]int{}, len(nums)
	backtrackV3(nums, []int{}, n, &res)
	return res
}

func backtrackV3(nums, track []int, n int, res *[][]int) {
	if len(track) == n {
		*res = append(*res, append([]int{}, track...))
		return
	}
	visited := map[int]bool{}
	j := len(track)
	for i := j; i < n; i++ {
		if visited[nums[i]] {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		backtrackV3(nums, append(track, nums[j]), n, res)
		nums[i], nums[j] = nums[j], nums[i]
		visited[nums[i]] = true
	}
}
