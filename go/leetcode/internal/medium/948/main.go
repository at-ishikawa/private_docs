package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})

	maxScore := 0
	score := 0
	left := 0
	right := len(tokens) - 1
	for left <= right {
		if tokens[left] <= power {
			score++
			if score > maxScore {
				maxScore = score
			}
			power -= tokens[left]
			left++
			continue
		}
		if score > 0 {
			power += tokens[right]
			right--
			score--
			continue
		}
		break
	}
	return maxScore
}
