package assign_cookies

import "sort"

// https://leetcode.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	ng, ns := len(g), len(s)
	if ng == 0 || ns == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i, j, res := 0, 0, 0
	for i < ng && j < ns {
		if g[i] <= s[j] {
			i++
			res++
		}
		j++
	}
	return res
}

func findContentChildrenV2(g []int, s []int) int {
	ng, ns := len(g), len(s)
	if ng == 0 || ns == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for ; i < ns && j < ng; i++ {
		if g[j] <= s[i] {
			j++
		}
	}
	return j
}
