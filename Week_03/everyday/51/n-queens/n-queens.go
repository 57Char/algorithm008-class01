package n_queens

// https://leetcode.com/problems/n-queens/

// 输入棋盘边长 n，返回所有合法的放置
func solveNQueens(n int) [][]string {
	res := [][]string{}

	// '.' 表示空，'Q' 表示皇后，初始化空棋盘
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	backtrack(board, n, 0, &res)

	return res
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack(board [][]rune, n, row int, res *[][]string) {
	if row == n {
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			arr[i] = string(board[i])
		}
		*res = append(*res, arr)
		return
	}

	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, n, row, col) {
			continue
		}
		// 做选择
		board[row][col] = 'Q'
		// 进入下一行决策
		backtrack(board, n, row+1, res)
		// 撤销选择
		board[row][col] = '.'
	}
}

// 是否可以在 board[row][col] 放置皇后？
func isValid(board [][]rune, n, row, col int) bool {
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	i, j := row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	// 检查左上方是否有皇后互相冲突
	i, j = row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}
