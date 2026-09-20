package internal

import (
	"strings"
)

// N=len(words), M=words[i].length
// Time complexity: O(N*M + N*M) = O(NM)
// Space complexity: O(N*M + N*M) = O(NM)
func fullJustify(words []string, maxWidth int) []string {
	return fullJustify1(words, maxWidth)
}

func fullJustify1(words []string, maxWidth int) []string {
	lines := make([]string, 0)
	line := words[0]
	for i := 1; i < len(words); i++ {
		word := words[i]
		if len(line+" "+word) <= maxWidth {
			line = line + " " + word
			continue
		}
		lines = append(lines, line)
		line = word
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}

	// 1. There are multiple words in a line
	// 2. There is only one word in a line
	// 3. The last line
	result := make([]string, len(lines))
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		lineLength := len(line)
		wordsInLine := strings.Split(line, " ")
		spacesInLine := len(wordsInLine) - 1
		if spacesInLine == 0 {
			result[i] = line + strings.Repeat(" ", maxWidth-lineLength)
			continue
		}
		extraSpaceCountInLine := (maxWidth - lineLength)

		justifiedLine := wordsInLine[0]

		oneMoreSpaceIndex := extraSpaceCountInLine % spacesInLine
		for j := 1; j < len(wordsInLine); j++ {
			// 5 / 3 => (2, 2, 1)
			extraSpace := extraSpaceCountInLine / spacesInLine
			if j <= oneMoreSpaceIndex {
				extraSpace++
			}
			justifiedLine = justifiedLine + strings.Repeat(" ", extraSpace+1)
			justifiedLine += wordsInLine[j]
		}
		result[i] = justifiedLine
	}
	endIndex := len(lines) - 1
	result[endIndex] = lines[endIndex] + strings.Repeat(" ", maxWidth-len(lines[endIndex]))

	return result
}
