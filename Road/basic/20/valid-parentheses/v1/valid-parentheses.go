package v1

// https://leetcode.com/problems/valid-parentheses/
// 20. Valid Parentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, c1 := range s {
		if c2, ok := pairs[c1]; ok {
			if len(stack) == 0 {
				return false
			}
			c := stack[len(stack)-1]
			if c != c2 {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c1)
		}
	}
	return len(stack) == 0
}
