package number_of_segments_in_a_string

// https://leetcode.com/problems/number-of-segments-in-a-string/

func countSegments(s string) int {
	res := 0
	for i, c := range s {
		if (i == 0 || s[i-1] == ' ') && c != ' ' {
			res++
		}
	}
	return res
}
