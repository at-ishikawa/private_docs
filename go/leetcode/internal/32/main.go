package main

// )()())()()())(

// find a corresponding ) for (
// (())()
// ( <- starts from index 1
// (, ( <- starts from index 2

// openParenthesisQueue <- [0, 1]
// endIndex <- 2 => 2-1 => 1 = 1
// endIndex <- 3 => 3 - 0 => 3 + 1
// answer <-

// https://leetcode.com/problems/longest-valid-parentheses/solutions/1139990/longest-valid-parentheses-short-easy-w-explanation-using-stack/
// Time complex: O(n)
// Space complexity: O(n)
func longestValidParentheses1(s string) int {
	stack := make([]int, 0)
	answer := 0
	left := 0

	for index, ch := range s {
		if ch == '(' {
			stack = append(stack, index)
			continue
		}
		// ch == ')'
		if len(stack) <= 0 {
			// no corresponding (.
			// Use the index for the current ) in the future as the starting point of valid parenthesis
			left = index + 1
			continue
		}

		// found a corresponding (. Simply pop the index
		// and calculate the length
		stack = stack[:len(stack)-1]

		currentLength := index - left + 1
		if len(stack) > 0 {
			// there are still (, which can be closed
			// get the length before this open parenthesis
			currentLength = index - stack[len(stack)-1]
		}
		if answer < currentLength {
			answer = currentLength
		}

	}

	return answer
}

// dp
func longestValidParentheses2(s string) int {
	// dp[i] = the valid parenthesis at the index i
	// if dp[i] = '(', it's 0
	// if dp[i] = ')',
	//   and s[i-1] == '(', then it's 2
	//   and s[i-1] == ')', then it's s[j] (where j = (i-1)-dp[i-1]) == '(', (finding a corresponding open parethensis), then, dp[i] = dp[i-1] + dp[j] + 2

	// )()())()
	// dp [0, 0, 2, 0, 4?, 0, 0, 2]
	// dp[4] = s[4] == ')' and s[3] == '(', and dp[2] = 2 => +2?

	// ((()))()
	// dp: [0, 0, 0, 2, 4, 6, 0, 8]
	dp := make([]int, len(s))
	for index, ch := range s {
		if ch == '(' {
			continue
		}

		if index == 0 {
			continue
		}
		if s[index-1] == '(' {
			// finding a corresponding (.
			// Add the previous longest length + 2
			dp[index] = 2
			if index >= 2 {
				dp[index] += dp[index-2]
			}
			continue
		}
		openIndex := (index - 1) - dp[index-1]
		if s[openIndex] != '(' {
			// no corresponding (. Invalid )
			continue
		}
		dp[index] = dp[index-1] + 2
		if openIndex > 1 {
			// add the length just before the open parenthesis
			dp[index] += dp[openIndex-1]
		}
	}

	answer := 0
	for _, length := range dp {
		if answer < length {
			answer = length
		}
	}

	return answer
}

// Wrong answer
// Time complexity: O(N)
// Space complexity: O(1)
// func longestValidParentheses(s string) int {
// 	answer := 0
// 	parenthesisCounter := 0
// 	currentAnswer := 0
// 	for _, ch := range s {
// 		if ch == '(' {
// 			parenthesisCounter++
// 			continue
// 		}
// 		if ch == ')' {
// 			if parenthesisCounter > 0 {
// 				parenthesisCounter--
// 				currentAnswer++
// 				if currentAnswer > answer {
// 					answer = currentAnswer
// 				}
// 				continue
// 			}
// 			// invalid parenthesis.
// 			currentAnswer = 0
// 		}
// 	}
// 	return 2 * answer
// }
