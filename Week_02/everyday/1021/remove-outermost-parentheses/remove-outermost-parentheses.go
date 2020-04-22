package remove_outermost_parentheses

import "strings"

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

func removeOuterParenthesesV2(S string) string {
	var counter int
	var sb strings.Builder
	sb.Grow(40)
	for _, c := range S {
		if c == '(' {
			if counter > 0 {
				sb.WriteRune(c)
			}
			counter++
		} else {
			counter--
			if counter > 0 {
				sb.WriteRune(c)
			}
		}
	}
	return sb.String()
}

func removeOuterParenthesesV3(S string) string {
	count, res := 0, []rune{}
	for _, c := range S {
		if c == '(' {
			if count > 0 {
				res = append(res, c)
			}
			count++
		} else {
			count--
			if count > 0 {
				res = append(res, c)
			}
		}
	}
	return string(res)
}

// very slow
func removeOuterParenthesesV4(S string) string {
	count, res := 0, ""
	for _, c := range S {
		if c == '(' {
			if count > 0 {
				res += string(c)
			}
			count++
		} else {
			count--
			if count > 0 {
				res += string(c)
			}
		}
	}
	return res
}