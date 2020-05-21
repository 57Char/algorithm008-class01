package word_ladder

// https://leetcode.com/problems/word-ladder/

// so slow
func ladderLength(beginWord string, endWord string, wordList []string) int {
	words, visited := map[string]bool{}, map[string]bool{}
	for _, w := range wordList {
		words[w] = true
	}
	if !words[endWord] {
		return 0
	}
	level := 1
	queue := []string{beginWord}
	visited[beginWord] = true

	for len(queue) > 0 {
		size := len(queue)
		for ; size > 0; size-- {
			lastWord := queue[0]
			queue = queue[1:]
			for i := 0; i < len(lastWord); i++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := lastWord[:i] + string(j) + lastWord[i+1:]
					if !words[newWord] || newWord == lastWord {
						continue
					}
					if newWord == endWord {
						return level + 1
					}
					if !visited[newWord] {
						queue = append(queue, newWord)
						visited[newWord] = true
					}
				}
			}
		}
		level++
	}

	return 0
}
