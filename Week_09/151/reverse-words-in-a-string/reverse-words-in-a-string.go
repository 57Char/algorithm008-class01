package reverse_words_in_a_string

import "strings"

// https://leetcode.com/problems/reverse-words-in-a-string/

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
func reverseWords(s string) string {
	list := strings.Split(strings.TrimSpace(s), " ")
	res := []string{}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == "" {
			continue
		}
		res = append(res, list[i])
	}
	return strings.Join(res, " ")
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
func reverseWordsV2(s string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	w := []string{}
	i, j := 0, 0
	for {
		for i < n && s[i] == ' ' {
			i++
		}
		if i >= n {
			break
		}
		j = i + 1
		for j < n && s[j] != ' ' {
			j++
		}
		w = append(w, s[i:j])
		i = j + 1
	}
	reverse(w, 0, len(w)-1)
	return strings.Join(w, " ")
}

func reverse(s []string, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.5 MB , 在所有 Go 提交中击败了 100.00% 的用户
func reverseWordsV3(s string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	w := []string{}
	i, j := 0, 0
	for {
		for i < n && s[i] == ' ' {
			i++
		}
		if i >= n {
			break
		}
		j = i + 1
		for j < n && s[j] != ' ' {
			j++
		}
		w = append(w, s[i:j])
		i = j + 1
	}
	n = len(w)
	for i := 0; i < n/2; i++ {
		w[i], w[n-i-1] = w[n-i-1], w[i]
	}
	return strings.Join(w, " ")
}

// slow
func reverseWordsV4(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	s = reverseV4(s, 0, len(s)-1)
	w := strings.Split(s, " ")
	res := []string{}
	for _, v := range w {
		if v != "" {
			res = append(res, reverseV4(v, 0, len(v)-1))
		}
	}
	return strings.Join(res, " ")
}

func reverseV4(s string, l, r int) string {
	c := []rune(s)
	for l < r {
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}
	return string(c)
}