package v1

// https://leetcode.com/problems/two-sum/
// 1. Two Sum

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	visited := map[int]int{}
	for i, v := range nums {
		if j, ok := visited[target-v]; ok {
			return []int{j, i}
		}
		visited[v] = i
	}
	return []int{}
}
