package v1

// https://leetcode.com/problems/valid-parentheses/
// 20. Valid Parentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || pairs[c] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

