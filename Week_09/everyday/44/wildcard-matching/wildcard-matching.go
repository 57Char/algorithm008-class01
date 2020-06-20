package wildcard_matching

// https://leetcode.com/problems/wildcard-matching/

func isMatch(s string, p string) bool {
	prev, curr := make([]bool, len(p)+1), make([]bool, len(p)+1)
	for i := 0; i <= len(s); i++ {
		curr, prev = prev, curr
		curr[0] = i == 0
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				curr[j] = prev[j] || prev[j-1] || curr[j-1]
			} else {
				curr[j] = prev[j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}
	return curr[len(p)]
}
