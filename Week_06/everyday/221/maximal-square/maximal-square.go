package maximal_square

// https://leetcode.com/problems/maximal-square/

func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	if col == 0 {
		return 0
	}
	dp := make([]int, col+1)

	side, prev := 0, 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			next := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(min(dp[j-1], dp[j]), prev) + 1
				side = max(side, dp[j])
			} else {
				dp[j] = 0
			}
			prev = next
		}
	}
	return side * side
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
理解 min(上, 左, 左上) + 1

// 伪代码
if (matrix(i - 1, j - 1) == '1') {
	dp(i, j) = min(dp(i - 1, j), dp(i, j - 1), dp(i - 1, j - 1)) + 1;
}

其中，dp(i, j) 是以 matrix(i - 1, j - 1) 为 右下角 的正方形的最大边长。
等同于：dp(i + 1, j + 1) 是以 matrix(i, j) 为右下角的正方形的最大边长

翻译成中文
- 若某格子值为 1，则以此为右下角的正方形的、最大边长为：上面的正方形、左面的正方形或左上的正方形中，最小的那个，再加上此格。
*/


