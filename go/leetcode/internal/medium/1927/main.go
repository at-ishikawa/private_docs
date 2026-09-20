// https://leetcode.com/problems/sum-game/description/
package main

// ??1199??
func sumGame(num string) bool {
	parseNum := func(num string) (int, int) {
		sum := 0
		questionCount := 0
		for _, digit := range num {
			if digit == '?' {
				questionCount++
				continue
			}
			sum += int(digit - '0')
		}
		return sum, questionCount
	}
	abs := func(val int) int {
		if val < 0 {
			return -val
		}
		return val
	}

	firstSum, firstCount := parseNum(num[:len(num)/2])
	secondSum, secondCount := parseNum(num[len(num)/2:])
	if abs(firstSum-secondSum) >= 10 {
		// Alice will win
		return true
	}
}
