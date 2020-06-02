package friend_circles

// https://leetcode.com/problems/friend-circles/

// Runtime: 16 ms, faster than 100.00% of Go online submissions for Friend Circles.
// Memory Usage: 6.1 MB, less than 100.00% of Go online submissions for Friend Circles.
func findCircleNum(M [][]int) int {
	n := len(M)
	u := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}
	return u.Count()
}



