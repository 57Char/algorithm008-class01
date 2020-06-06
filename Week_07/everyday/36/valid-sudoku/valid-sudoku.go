package valid_sudoku

// https://leetcode.com/problems/valid-sudoku/

// 执行用时 : 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 2.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
func isValidSudoku(board [][]byte) bool {
	var bitmap [3][9][9]int
	for i, line := range board {
		for j, cell := range line {
			if cell == '.' {
				continue
			}
			k, index := (i/3)*3+j/3, cell-'1'
			if bitmap[0][i][index] == 1 || bitmap[1][j][index] == 1 || bitmap[2][k][index] == 1 {
				return false
			}
			bitmap[0][i][index], bitmap[1][j][index], bitmap[2][k][index] = 1, 1, 1
		}
	}
	return true
}

// 执行用时 : 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 2.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
func isValidSudokuV2(board [][]byte) bool {
	var row, col, box [9][9]int
	for i, line := range board {
		for j, cell := range line {
			if cell == '.' {
				continue
			}
			k, index := (i/3)*3+j/3, cell-'1'
			if row[i][index] == 1 || col[j][index] == 1 || box[k][index] == 1 {
				return false
			}
			row[i][index], col[j][index], box[k][index] = 1, 1, 1
		}
	}
	return true
}

// 执行用时 : 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 4.1 MB, 在所有 Go 提交中击败了 9.52% 的用户
func isValidSudokuV3(board [][]byte) bool {
	var rowSet, colSet, boxSet []map[byte]bool
	for i := 0; i < 9; i++ {
		rowSet = append(rowSet, map[byte]bool{})
		colSet = append(colSet, map[byte]bool{})
		boxSet = append(boxSet, map[byte]bool{})
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			k := (i/3)*3 + j/3
			if rowSet[i][val] || colSet[j][val] || boxSet[k][val] {
				return false
			}
			rowSet[i][val], colSet[j][val], boxSet[k][val] = true, true, true
		}
	}
	return true
}
