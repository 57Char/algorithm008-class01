package isomorphic_strings

import "strings"

// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	mem := make([]byte, 1<<7, 1<<7)
	uniq := make([]byte, 1<<7, 1<<7)
	for i := range s {
		if mem[s[i]] == 0 && uniq[t[i]] == 0 {
			mem[s[i]] = t[i]
			uniq[t[i]] = s[i]
		}
		if mem[s[i]] != t[i] {
			return false
		}
	}
	return true
}

// slow
func isIsomorphicV2(s string, t string) bool {
	sc, tc := map[byte]int{}, map[byte]int{}
	for i := range s {
		if sc[s[i]] != tc[t[i]] {
			return false
		}
		sc[s[i]] = i + 1
		tc[t[i]] = i + 1
	}
	return true
}

func isIsomorphicV3(s string, t string) bool {
	sc, tc := make([]int, 1<<7), make([]int, 1<<7)
	for i := range s {
		si, ti := strings.IndexByte(s, s[i]), strings.IndexByte(t, t[i])
		if sc[si] == 0 {
			sc[si] = i + 1
		}
		if tc[ti] == 0 {
			tc[ti] = i + 1
		}
		if sc[si] != tc[ti] {
			return false
		}
	}
	return true
}

func isIsomorphicV4(s string, t string) bool {
	sc, tc := make([]int, 1<<7), make([]int, 1<<7)
	for i := range s {
		si, ti := strings.IndexByte(s, s[i]), strings.IndexByte(t, t[i])
		if sc[si] != tc[ti] {
			return false
		}
		sc[si] = i + 1
		tc[ti] = i + 1
	}
	return true
}
