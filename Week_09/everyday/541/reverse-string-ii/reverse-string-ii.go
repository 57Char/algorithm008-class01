package reverse_string_ii

// https://leetcode.com/problems/reverse-string-ii/

// 执行用时 : 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 3.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
func reverseStr(s string, k int) string {
	c, n := []rune(s), len(s)

	for i := 0; i < n; i += 2 * k {
		j := i + k - 1
		if j > n-1 {
			j = n - 1
		}
		reverse(c, i, j)
	}
	return string(c)
}

func reverse(c []rune, l, r int) {
	for l < r {
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}
}