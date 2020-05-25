package house_robber

// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	prev, curr := 0, 0
	for _, v := range nums {
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

func robV2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}
