// https://leetcode.com/problems/guess-the-word/?envType=study-plan-v2&envId=google-spring-23-high-frequency

package main

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */

type Master struct {
}

func (this *Master) Guess(word string) int {
	return 0
}

// pass a word => we know the number of words matching
// based on the number of matching, we can use the another word

// secret: "abcdef", words: ["abcdee","abcdeg","abcdef", "abccee", "abbdeg"]
// the 0th word "abcdee" and guessed. Got 5
// Check other words that are 5 characters matching and got "abcdeg","abcdef", "abccee"
// Use the first candidate of the word "abcdeg", then got 5
// From "abcdef", "abccee", check matching 5 from "abcdeg", and got "abcdef"
// Use "abcdef", then got a word 6 word matching

// Time complexity: O(N * Number of guesses)
// Space complexity: O(N)
func countMatchLetters(str1 string, str2 string) int {
	countMatchLetters := 0
	for j := 0; j < len(str1); j++ {
		if str1[j] == str2[j] {
			countMatchLetters++
		}
	}
	return countMatchLetters
}

func findSecretWord(words []string, master *Master) {
	getNextCandidates := func(baseWord string, candidates []string, expectedCountLetters int) []string {
		result := make([]string, 0)
		for i := 0; i < len(candidates); i++ {
			if countMatchLetters(candidates[i], baseWord) == expectedCountLetters {
				result = append(result, candidates[i])
			}
		}
		return result
	}

	matchLetterMap := make(map[string]map[string]int, 0)
	for _, word := range words {
		matchLetterMap[word] = make(map[string]int, 0)
	}
	for _, word := range words {
		for _, word2 := range words {
			count := countMatchLetters(word, word2)
			matchLetterMap[word][word2] = count
		}
	}

	getHighestScoreWord := func(words []string) string {
		highestScore := 0
		resultWord := ""
		for _, word := range words {
			score := 0

		INNER:
			for mapWord, count := range matchLetterMap[word] {
				// todo: Optimize to remove a word from a mapping of a matching letter count
				for _, word := range words {
					if word == mapWord {
						score += count
						continue INNER
					}
				}
			}
			if highestScore < score {
				highestScore = score
				resultWord = word
			}
		}
		return resultWord
	}

	for {
		word := getHighestScoreWord(words)
		countMatchLetters := master.Guess(word)
		if countMatchLetters == 6 {
			// found the secret!
			break
		}

		words = getNextCandidates(word, words, countMatchLetters)
	}
}
