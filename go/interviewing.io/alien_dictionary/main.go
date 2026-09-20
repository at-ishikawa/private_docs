// https://interviewing.io/questions/alien-dictionary

package main

import (
	"fmt"
	"strings"
)

// Input: words = ["kaa","akcd","akca","cak","cad"]
// Output: "kdac

// [kaa, akcd] => k < a
// [kaa, akcd, akca] => d < a (from akc(d|a))
// [kaa, akcd, akca, cak] => a < c
// [kaa, akcd, akca, cak, cad] => k < d
// k < d < a < c

// Check the previous elements, and check the different letters, and put the order?
// Data:
// [kaa]
// [akcd] => [k, a]
// [akca] => [k, d, a] (not sure k < d)
// [cak] => [k, a, c], [d, a]
// [cad] => [k, d, a, c]
// [k, a], [d, a], [a, c], [k, d]
// [k, [a, d]], [a, c], [d, [a]]
// => kac, kdac

// Data:
// Input: words = ["b","a"]
// Output: "ba"
// Example 3
// Input: words = ["ab","a","b"]
// Output: ""

// Data: [ab, a, b]
// [ab, a]: ?
// [ab, a, b]: a < b

func getAlianSubstr(allDependencies map[rune][]rune, ch rune, aliasSubstr map[rune]string) string {
	if result, ok := aliasSubstr[ch]; ok {
		return result
	}

	substringSet := make([]string, 0)
	for _, dependency := range allDependencies[ch] {
		substr := getAlianSubstr(allDependencies, dependency, aliasSubstr)

		isFound := false
		for _, existingSet := range substringSet {
			if strings.Contains(existingSet, substr) {
				isFound = true
			}
		}
		if isFound {
			continue
		}

		// concatenate independent substrings
		isOverwritten := false
		for i := 0; i < len(substringSet); i++ {
			if strings.Contains(substr, substringSet[i]) {
				substringSet[i] = substr
				isOverwritten = true
				break
			}
		}
		if !isOverwritten {
			substringSet = append(substringSet, substr)
		}
	}

	result := string(ch) + strings.Join(substringSet, "")
	aliasSubstr[ch] = result
	return result
}

func sortAlianWords(sortedAliendWords []string) string {
	dependencies := make(map[rune][]rune, 0)
	for _, word := range sortedAliendWords {
		for _, ch := range word {
			if _, ok := dependencies[ch]; ok {
				continue
			}
			dependencies[ch] = make([]rune, 0)
		}
	}

	var emptyRune rune
	var firstCh rune
	for i := 1; i < len(sortedAliendWords); i++ {
		previousWord := sortedAliendWords[i-1]
		word := sortedAliendWords[i]

		for j := 0; j < len(word) && j < len(previousWord); j++ {
			previousCh := rune(previousWord[j])
			currentCh := rune(word[j])
			if previousCh == currentCh {
				continue
			}

			if firstCh == emptyRune {
				firstCh = previousCh
			}

			dependencies[previousCh] = append(dependencies[previousCh], currentCh)
			break
		}
	}

	aliases := make(map[rune]string, 0)
	return getAlianSubstr(dependencies, firstCh, aliases)
}

func sortAlianWordsBFS(words []string) string {
	indegree := make(map[rune]int, 0)
	adjacents := make(map[rune][]byte, 0)
	for _, word := range words {
		for _, ch := range word {
			adjacents[ch] = make([]byte, 0)
			indegree[ch] = 0
		}
	}

	for i := 0; i < len(words)-1; i++ {
		firstWord := words[i]
		secondWord := words[i+1]
	}
}

func main() {
	testCases := []struct {
		name             string
		sortedAlianWords []string
		want             string
	}{
		{
			name: "example 1",
			sortedAlianWords: []string{
				"kaa", "akcd", "akca", "cak", "cad",
			},
			want: "kdac",
		},
		{
			name: "example 2",
			sortedAlianWords: []string{
				"b", "a",
			},
			want: "ba",
		},
		{
			name: "example 3",
			sortedAlianWords: []string{
				"ab", "a", "b",
			},
			want: "",
		},
	}

	for _, tc := range testCases {
		got := sortAlianWords(tc.sortedAlianWords)
		fmt.Printf("Test case %s: Want: %s, Got: %s\n", tc.name, tc.want, got)
	}
}
