package surrounded_regions

// https://leetcode.com/problems/surrounded-regions/

func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	if col == 0 {
		return
	}

	dummy := row * col
	u := NewUnionFind(dummy)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == row-1 || j == 0 || j == col-1 {
					u.Union(i*col+j, dummy)
				} else {
					for _, v := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
						x, y := v[0], v[1]
						if board[i+x][j+y] == 'O' {
							u.Union(i*col+j, (i+x)*col+(j+y))
						}
					}
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if u.Find(dummy) == u.Find(i*col+j) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
