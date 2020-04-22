package remove_outermost_parentheses

// https://leetcode.com/problems/remove-outermost-parentheses/

func removeOuterParentheses(S string) string {
	var res []rune
	counter := 0
	for _, c := range S {
		if c == ')' {
			counter--
		}
		if counter >= 1 {
			res = append(res, c)
		}
		if c == '(' {
			counter++
		}
	}
	return string(res)
}

