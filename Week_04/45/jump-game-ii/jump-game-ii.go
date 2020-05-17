package jump_game_ii

// https://leetcode.com/problems/jump-game-ii/

func jump(nums []int) int {
	res, end, maxPosition, n := 0, 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			res++
		}
	}
	return res
}

func jumpV2(nums []int) int {
	res, end, maxPosition, n := 0, 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			res++
		}
		if end >= n {
			break
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