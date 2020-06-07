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

func lengthOfLISV3(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	tails := make([]int, n)
	tails[0] = nums[0]
	end := 0
	for i := 1; i < n; i++ {
		num := nums[i]
		// 比 tail 数组实际有效的末尾的那个元素还大
		if num > tails[end] {
			end++
			tails[end] = num
			continue
		}
		// 找到第 1 个大于等于 nums[i] 的元素，尝试让那个元素更小
		lo, hi := 0, end
		for lo < hi {
			// 选左中位数
			mi := lo + (hi-lo)>>1
			if tails[mi] < num {
				// 中位数肯定不是要找的数
				lo = mi + 1
			} else {
				hi = mi
			}
		}
		tails[lo] = num
	}
	return end + 1
}

// https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/
//1、设置一个数组 tail，初始时为空；
//注意：数组 tail 虽然是有序数组，但它不是问题中的「最长上升子序列」（下文还会强调），不能命名为 LIS。有序数组 tail 只是用于求解 LIS 问题的辅助数组。
//2、在遍历数组 nums 的过程中，每来一个新数 num，如果这个数严格大于有序数组 tail 的最后一个元素，就把 num 放在有序数组 tail 的后面，否则进入第 3 点；
//注意：这里的大于是「严格大于」，不包括等于的情况。
//3、在有序数组 tail 中查找第 1 个等于大于 num 的那个数，试图让它变小；
//如果有序数组 tail 中存在等于 num 的元素，什么都不做，因为以 num 结尾的最短的「上升子序列」已经存在；
//如果有序数组 tail 中存在大于 num 的元素，找到第 1 个，让它变小，这样我们就找到了一个结尾更小的相同长度的上升子序列。
//说明：我们再看一下数组 tail[i] 的定义：长度为 i + 1 的所有最长上升子序列的结尾的最小值。因此，在遍历的过程中，我们试图让一个大的值变小是合理的。
