package main

import "fmt"

func minimumWindow(s string, t string) string {
	tMap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	result := ""
	startIndex := 0
	tCounter := len(t)
	for endIndex := 0; endIndex < len(s); endIndex++ {
		end := s[endIndex]
		if _, ok := tMap[end]; ok {
			tMap[end]--
			tCounter--
		}
		for tCounter == 0 {
			start := s[startIndex]
			startIndex++
			if _, ok := tMap[start]; !ok {
				continue
			}

			tMap[start]++
			tCounter++
			if tMap[start] <= 0 {
				continue
			}

			substr := s[startIndex-1 : endIndex+1]
			if result == "" || len(substr) < len(result) {
				result = substr
			}
			break
		}
	}
	return result
}

func main() {
	testCases := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2",
			s:    "ab",
			t:    "b",
			want: "b",
		},
		{
			name: "example 3",
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for _, tc := range testCases {
		got := minimumWindow(tc.s, tc.t)
		fmt.Printf("Test %s: Want %s, Got %s\n", tc.name, tc.want, got)
	}
}
