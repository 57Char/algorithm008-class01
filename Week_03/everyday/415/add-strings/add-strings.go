package add_strings

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	if n1 == 0 {
		return num2
	}
	if n2 == 0 {
		return num1
	}
	i, j, carry, res := n1-1, n2-1, 0, []string{}
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		res = append(res, strconv.Itoa(sum%10))
	}
	i, j = 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return strings.Join(res, "")
}

// very slow
func addStringsV2(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	if n1 == 0 {
		return num2
	}
	if n2 == 0 {
		return num1
	}
	i, j, carry, res := n1-1, n2-1, 0, ""
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	return res
}
