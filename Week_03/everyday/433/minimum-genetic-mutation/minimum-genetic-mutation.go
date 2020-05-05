package minimum_genetic_mutation

import "math"

// https://leetcode.com/problems/minimum-genetic-mutation/

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	visited, banks := map[string]bool{}, map[string]bool{}
	for _, b := range bank {
		banks[b] = true
	}
	genes := []byte{'A', 'C', 'G', 'T'}

	level := 0
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		n := len(queue)
		for ; n > 0; n-- {
			curr := queue[0]
			queue = queue[1:]
			if curr == end {
				return level
			}
			newStart := []byte(curr)
			for i, v := range newStart {
				for _, g := range genes {
					newStart[i] = g
					next := string(newStart)
					if !visited[next] && banks[next] {
						queue = append(queue, next)
						visited[next] = true
					}
				}
				newStart[i] = v
			}
		}
		level++
	}

	return -1
}

func minMutationV2(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	visited := map[string]bool{}
	level := 0
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		size := len(queue)
		for ; size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]
			if curr == end {
				return level
			}
			for _, v := range bank {
				if !visited[v] && isValid(curr, v) {
					queue = append(queue, v)
					visited[v] = true
				}
			}
		}
		level++
	}

	return -1
}

func isValid(s1, s2 string) bool {
	diff := false
	for i, v := range []byte(s1) {
		if v != s2[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return diff
}

func minMutationV3(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	level, n := 0, len(bank)
	visited := make([]bool, n)
	queue := []string{start}

	for len(queue) > 0 {
		size := len(queue)
		for ; size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]
			if curr == end {
				return level
			}
			for i, v := range bank {
				if !visited[i] && isValid(curr, v) {
					queue = append(queue, v)
					visited[i] = true
				}
			}
		}
		level++
	}

	return -1
}

func minMutationV4(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	res := math.MaxInt64
	dfs(start, end, bank, make([]bool, len(bank)), 0, &res)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func dfs(start, end string, bank []string, visited []bool, level int, res *int) {
	if start == end {
		if *res > level {
			*res = level
		}
		return
	}
	for i, v := range bank {
		if !visited[i] && isValid(start, v) {
			visited[i] = true
			dfs(v, end, bank, visited, level+1, res)
			visited[i] = false
		}
	}
}
