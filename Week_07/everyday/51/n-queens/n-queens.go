package n_queens

// https://leetcode.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	res := [][]string{}
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
		if !isValid(board, n, row, col) {
			continue
		}
		board[row][col] = 'Q'
		backtrack(board, n, row+1, res)
		board[row][col] = '.'
	}
}

func isValid(board [][]rune, n, row, col int) bool {
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
