package ti_huan_kong_ge_lcof

import "strings"

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

func replaceSpace(s string) string {
	var res []rune
	for _, v := range s {
		if v == ' ' {
			res = append(res, []rune("%20")...)
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}

func replaceSpaceV2(s string) string {
	var res []rune
	for _, v := range s {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}

func replaceSpaceV3(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
