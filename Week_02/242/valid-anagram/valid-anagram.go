package valid_anagram

// https://leetcode.com/problems/valid-anagram/

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
		if v > 0 {
			return false
		}
	}
	return true
}