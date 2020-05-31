package house_robber_ii

// https://leetcode.com/problems/house-robber-ii/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for House Robber II.
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(helper(nums[:n-1]), helper(nums[1:n]))
}

func helper(nums []int) int {
	prev, curr := 0, 0
	for _, v := range nums {
		// prev = dp[i-2]
		// curr = dp[i-1]
		prev, curr = curr, max(curr, prev+v)
	}
	return curr
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
