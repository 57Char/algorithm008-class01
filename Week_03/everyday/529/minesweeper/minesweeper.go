package minesweeper

// https://leetcode.com/problems/minesweeper/

func updateBoard(board [][]byte, click []int) [][]byte {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	row, col := click[0], click[1]
	if board[row][col] == 'M' || board[row][col] == 'X' {
		board[row][col] = 'X'
		return board
	}
	m, n := len(board), len(board[0])
	var num byte
	for _, dir := range dirs {
		newRow := dir[0] + row
		newCol := dir[1] + col
		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && board[newRow][newCol] == 'M' {
			num++
		}
	}
	if num > 0 {
		board[row][col] = num + '0'
		return board
	}
	board[row][col] = 'B'
	for _, dir := range dirs {
		newRow := dir[0] + row
		newCol := dir[1] + col
		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && board[newRow][newCol] == 'E' {
			updateBoard(board, []int{newRow, newCol})
		}
	}
	return board
}

func updateBoardV2(board [][]byte, click []int) [][]byte {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	m, n := len(board), len(board[0])
	helper(&board, click, m, n, dirs)
	return board
}

func helper(board *[][]byte, click []int, m, n int, dirs [][]int) {
	row, col := click[0], click[1]
	if (*board)[row][col] == 'M' || (*board)[row][col] == 'X' {
		(*board)[row][col] = 'X'
		return
	}
	var num byte
	for _, dir := range dirs {
		newRow := dir[0] + row
		newCol := dir[1] + col
		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && (*board)[newRow][newCol] == 'M' {
			num++
		}
	}
	if num > 0 {
		(*board)[row][col] = num + '0'
		return
	}
	(*board)[row][col] = 'B'
	for _, dir := range dirs {
		newRow := dir[0] + row
		newCol := dir[1] + col
		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && (*board)[newRow][newCol] == 'E' {
			helper(board, []int{newRow, newCol}, m, n, dirs)
		}
	}
}
