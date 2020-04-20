package valid_anagram

// https://leetcode.com/problems/valid-anagram/

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	m := make([]int, 26)
	for i := 0; i < n; i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramV3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for i, v := range s {
		m[v]++
		m[rune(t[i])]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// very slow
func isAnagramV4(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, t1 := strings.Split(s, ""), strings.Split(t, "")
	sort.Strings(s1)
	sort.Strings(t1)
	s, t = strings.Join(s1, ""), strings.Join(t1, "")
	return s == t
}
