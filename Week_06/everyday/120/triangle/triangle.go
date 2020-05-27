package triangle

// https://leetcode.com/problems/triangle/

// Runtime: 4 ms, faster than 95.54% of Go online submissions for Triangle.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Triangle.
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}
	for i := n - 2; i >= 0; i-- {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
