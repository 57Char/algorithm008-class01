package jump_game_ii

// https://leetcode.com/problems/jump-game-ii/

func jump(nums []int) int {
	res, end, step, n := 0, 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		step = max(step, i+nums[i])
		if i == end {
			end = step
			res++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}