package word_ladder

// https://leetcode.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := map[string]bool{}
	for _, word := range wordList {
		words[word] = true
	}
	if !words[endWord] {
		return 0
	}
	beginWords, endWords := map[string]bool{}, map[string]bool{}
	beginWords[beginWord] = true
	endWords[endWord] = true
	res := 1
	for len(beginWords) > 0 {
		res++
		newWords := map[string]bool{}
		for word := range beginWords {
			for i, char := range word {
				for c := 'a'; c <= 'z'; c++ {
					if c != char {
						newWord := word[:i] + string(c) + word[i+1:]
						if endWords[newWord] {
							return res
						}
						if words[newWord] {
							newWords[newWord] = true
							delete(words, newWord)
						}
					}
				}
			}
		}
		beginWords = newWords
		if len(beginWords) > len(endWords) {
			beginWords, endWords = endWords, beginWords
		}
	}
	return 0
}
