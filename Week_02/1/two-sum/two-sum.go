package two_sum

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	visited := map[int]int{}
	for i, val := range nums {
		if j, ok := visited[target-val]; ok {
			return []int{j, i}
		}
		visited[val] = i
	}
	return []int{}
}
