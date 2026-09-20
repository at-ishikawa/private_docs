package internal

// endWord -> s_i -> .. -> beginWord
// minSteps: map[string]int
// bfs: minSteps

type queueItem struct {
	word  string
	count int
}

func getLetterDiffCount(word1 string, word2 string) int {
	count := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			count++
		}
	}
	return count
}

// Time complexity: O(N * N) => O(N^2)
// Space complexity: O(N + N) => O(N)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	/*
		adjacentWordMap := make(map[string][]string)
		for _, word := range wordList {
			adjacentWords := make([]string, 0)
			for _, otherWord := range wordList {
				if diffCount := getLetterDiffCount(word, otherWord); diffCount != 1 {
					continue
				}
				adjacentWords = append(adjacentWords, otherWord)
			}
			adjacentWordMap[word] = adjacentWords
		}
	*/

	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{
		word:  beginWord,
		count: 1,
	})
	checked := make(map[string]bool, len(wordList))

	answer := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.word == endWord {
			answer = item.count
			break
		}

		if _, ok := checked[item.word]; ok {
			continue
		}

		for _, word := range wordList {
			diffCount := getLetterDiffCount(word, item.word)
			if diffCount == 1 {
				queue = append(queue, queueItem{
					word:  word,
					count: item.count + 1,
				})
			}
		}
		checked[item.word] = true
	}
	return answer
}
