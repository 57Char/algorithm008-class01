package bulls_and_cows

import "fmt"

// https://leetcode.com/problems/bulls-and-cows/

func getHint(secret string, guess string) string {
	counter := make([]int, 10)
	bulls, cows, n := 0, 0, len(secret)
	for i := 0; i < n; i++ {
		if v1, v2 := secret[i]-'0', guess[i]-'0'; v1 == v2 {
			bulls++
		} else {
			if counter[v1] < 0 {
				cows++
			}

			if counter[v2] > 0 {
				cows++
			}

			counter[v1]++
			counter[v2]--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHintV2(secret string, guess string) string {
	counter := make([]int, 10)
	s, g := []byte(secret), []byte(guess)
	bulls, cows := 0, 0
	for i, v := range s {
		if v == g[i] {
			bulls++
		} else {
			if counter[v-'0'] < 0 {
				cows++
			}

			if counter[g[i]-'0'] > 0 {
				cows++
			}

			counter[v-'0']++
			counter[g[i]-'0']--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHintV3(secret string, guess string) string {
	counter := make([]int, 10)
	s, g := []byte(secret), []byte(guess)
	bulls, cows := 0, 0
	for i, v := range s {
		if v == g[i] {
			bulls++
		} else {
			counter[v-'0']++
			if counter[v-'0'] <= 0 {
				cows++
			}

			counter[g[i]-'0']--
			if counter[g[i]-'0'] >= 0 {
				cows++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHintV4(secret string, guess string) string {
	m := map[byte]int{}
	s, g := []byte(secret), []byte(guess)
	bulls, cows := 0, 0
	for i, v := range s {
		if v == g[i] {
			bulls++
		} else {
			m[v]++
			if m[v] <= 0 {
				cows++
			}

			m[g[i]]--
			if m[g[i]] >= 0 {
				cows++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHintV5(secret string, guess string) string {
	c1, c2 := make([]int, 10), make([]int, 10)
	s, g := []byte(secret), []byte(guess)
	bulls, cows := 0, 0
	for i, v := range s {
		if v == g[i] {
			bulls++
		} else {
			c1[v-'0']++
			c2[g[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		v1, v2 := c1[i], c2[i]
		if v1 > v2 {
			cows += v2
		} else {
			cows += v1
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}