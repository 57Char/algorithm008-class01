package generate_parentheses

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}
	generate(0, 0, n, []byte{}, &res)
	return res
}

func generate(l, r, n int, s []byte, res *[]string) {
	if l == n && r == n {
		*res = append(*res, string(s))
		return
	}
	if l < n {
		generate(l+1, r, n, append(s, '('), res)
	}
	if r < l {
		generate(l, r+1, n, append(s, ')'), res)
	}
}

func generateParenthesisV2(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}
	generateV2(n, n, []byte{}, &res)
	return res
}

func generateV2(l, r int, s []byte, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, string(s))
		return
	}
	if l > 0 {
		generateV2(l-1, r, append(s, '('), res)
	}
	if r > l {
		generateV2(l, r-1, append(s, ')'), res)
	}
}