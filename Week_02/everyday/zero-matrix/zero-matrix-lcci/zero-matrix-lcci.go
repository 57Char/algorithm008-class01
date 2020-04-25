package zero_matrix_lcci

// https://leetcode-cn.com/problems/zero-matrix-lcci/

func setZeroes(matrix [][]int) {
	mi, mj := map[int]bool{}, map[int]bool{}
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				mi[i], mj[j] = true, true
			}
		}
	}
	for i, row := range matrix {
		for j := range row {
			if mi[i] || mj[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesV2(matrix [][]int) {
	nRow := len(matrix)
	if nRow == 0 {
		return
	}
	nCol := len(matrix[0])
	row, col := make([]int, 0, nRow), make([]int, 0, nCol)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}
	for _, i := range row {
		for j := 0; j < nCol; j++ {
			matrix[i][j] = 0
		}
	}
	for _, j := range col {
		for i := 0; i < nRow; i++ {
			matrix[i][j] = 0
		}
	}
}