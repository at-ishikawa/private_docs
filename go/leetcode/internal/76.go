package internal

import "fmt"

func minWindow_76(s string, t string) string {
	tCounts := make(map[byte]int)
	for _, tCh := range t {
		tCounts[byte(tCh)]++
	}

	startIndex := 0
	endIndex := 0
	tCounter := 0
	minimumSubstr := ""
	for endIndex < len(s) {
		sCh := s[endIndex]
		count, ok := tCounts[sCh]
		if ok {
			tCounts[sCh] = count - 1

			// we remove a counter only when there is a character no more than t's counts
			if count > 0 {
				tCounter++
			}
		}

		for ; startIndex <= endIndex && tCounter == len(t); startIndex++ {
			substr := s[startIndex : endIndex+1]
			fmt.Printf("(%d, %d): %s\n", startIndex, endIndex, substr)
			if minimumSubstr == "" || len(substr) < len(minimumSubstr) {
				minimumSubstr = substr
			}

			// move the start index
			startCh := s[startIndex]
			count, ok := tCounts[startCh]
			if ok {
				// we should count if the tCounts can be negative
				tCounts[startCh] = count + 1
				if count == 0 {
					tCounter--
				}
			}
		}
		endIndex++
	}
	return minimumSubstr
}
