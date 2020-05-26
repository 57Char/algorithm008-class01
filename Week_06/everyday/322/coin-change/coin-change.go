package coin_change

// https://leetcode.com/problems/coin-change/

// Runtime: 4 ms, faster than 99.47% of Go online submissions for Coin Change.
// Memory Usage: 6 MB, less than 100.00% of Go online submissions for Coin Change.
func coinChange(coins []int, amount int) int {
	if amount < 0 || len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i < c {
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}