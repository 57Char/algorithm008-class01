package search_a_2d_matrix

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n, m := len(matrix), len(matrix[0])

	if target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}

	mid, lo, hi := 0, 0, n-1
	for lo <= hi {
		mid = lo + (hi-lo)>>1
		val := matrix[mid][0]
		if target == val {
			return true
		} else if target < val {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	le, ri := 0, m-1
	for le <= ri {
		mid = le + (ri-le)>>1
		val := matrix[hi][mid]
		if target == val {
			return true
		} else if target < val {
			ri = mid - 1
		} else {
			le = mid + 1
		}
	}
	return false
}
