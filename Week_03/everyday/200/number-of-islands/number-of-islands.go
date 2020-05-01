package number_of_islands

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			res++
			dfs(grid, m, n, i, j)
		}
	}
	return res
}

func dfs(grid [][]byte, m, n, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, m, n, i-1, j)
	dfs(grid, m, n, i+1, j)
	dfs(grid, m, n, i, j-1)
	dfs(grid, m, n, i, j+1)
}

// TODO: 其他解法
