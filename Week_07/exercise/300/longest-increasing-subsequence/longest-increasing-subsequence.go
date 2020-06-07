package longest_increasing_subsequence

// https://leetcode.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 2.4 MB , 在所有 Go 提交中击败了 75.00% 的用户
func lengthOfLISV2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	/**
	dp[i]: 所有长度为i+1的递增子序列中, 最小的那个序列尾数.
	由定义知dp数组必然是一个递增数组, 可以用 res 来表示最长递增子序列的长度.
	对数组进行迭代, 依次判断每个数num将其插入dp数组相应的位置:
	1. num > dp[res], 表示num比所有已知递增序列的尾数都大, 将num添加入dp
	数组尾部, 并将最长递增序列长度res加1
	2. dp[i-1] < num <= dp[i], 只更新相应的dp[i]
	*/
	dp := make([]int, n)
	res := 0
	for _, num := range nums {
		lo, hi := 0, res
		for lo < hi {
			mi := lo + (hi-lo)>>1
			if dp[mi] < num {
				lo = mi + 1
			} else {
				hi = mi
			}
		}
		dp[lo] = num
		if res == hi {
			res++
		}
	}
	return res
}
