package distinct_subsequences

// https://leetcode.com/problems/distinct-subsequences/

func numDistinct(s string, t string) int {
	n1, n2 := len(s), len(t)
	dp := make([][]int, n1+1)
	// s不空，t空，次数为1
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][n2] = 1
	}
	// s空，t不空，次数为0
	//for i := 0; i <= n2; i++ {
	//	dp[0][i] = 0
	//}
	// 倒序，t每次增加1个字母
	for j := n2 - 1; j >= 0; j-- {
		// 倒序，s每次增加1哥字母
		for i := n1 - 1; i >= 0; i-- {
			if t[j] == s[i] {
				// 相等，选择当前字母，或，不选择当前字母
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				// 不相等，不选择当前字母
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
