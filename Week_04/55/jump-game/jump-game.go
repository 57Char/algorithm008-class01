package jump_game

// https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	target := n - 1
	for i := target; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}
	return target == 0
}

func canJumpV2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	farthest := 0
	for i := 0; i <= farthest; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if farthest >= n-1 {
			return true
		}
	}
	return false
}
