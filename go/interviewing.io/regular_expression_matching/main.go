package main

import "fmt"

func main() {
	testCases := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "example 1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "example 2",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "example 3",
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			name: "test case 1",
			s:    "acedegebgegddbedg",
			p:    "a.*b.*",
			want: true,
		},
		{
			name: "test case 2",
			s:    "aaaabdecdbc",
			p:    "a*b.*.*c",
			want: true,
		},
		{
			name: "test case 2",
			s:    "aabbbbbbbbbbbbbc",
			p:    "a.*bc",
			want: true,
		},
		{
			name: "test case 3",
			s:    "ab",
			p:    "a*b",
			want: true,
		},
	}

	for _, tc := range testCases {
		got := matchRegularExpression3(tc.s, tc.p)
		fmt.Printf("Test %s: ", tc.name)
		if tc.want != got {
			fmt.Printf("Failed. Want %v, Got %v", tc.want, got)
		} else {
			fmt.Printf("Passed")
		}
		fmt.Println()
	}
}

func matchRegularExpression3(s string, p string) bool {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(p))
	}
	if !(s[0] == p[0] || p[0] == '.') {
		return false
	}
	dp[0][0] = true

	// * can be 0 letter, so we have to mark it to a true
	for pIndex := 1; pIndex < len(p); pIndex++ {
		if p[pIndex] != '*' {
			break
		}
		dp[0][pIndex] = dp[0][pIndex-1]
	}

	for sIndex := 1; sIndex < len(s); sIndex++ {
		for pIndex := 1; pIndex < len(p); pIndex++ {
			switch p[pIndex] {
			case '*':
				if pIndex == 0 {
					return false
				}
				dp[sIndex][pIndex] = dp[sIndex-1][pIndex-1]
				previousChar := p[pIndex-1]
				if previousChar == '.' || previousChar == s[sIndex] {
					dp[sIndex][pIndex] = dp[sIndex][pIndex] || dp[sIndex-1][pIndex]
				}
			case '.':
				dp[sIndex][pIndex] = dp[sIndex-1][pIndex-1]
			default:
				dp[sIndex][pIndex] = dp[sIndex-1][pIndex-1] && s[sIndex] == p[pIndex]
			}
		}
	}

	return dp[len(s)-1][len(p)-1]
}

func matchRegularExpression2(s string, p string) bool {
	// dynamic programming
	// s=aaabcdd, p=a*b.*d
	// dp:  a     *   b . *
	//  a [true, -, -]
	//  a [-, true, -]
	//  a [-, true, false, -]
	//  b [-, false, true, -]
	//  c [-, -, false, true]
	//  d

	// [i, j]
	// ch != . nor *: check dp[i, j] = s[i] == p[j]
	//     if dp[i - 1, j - 1] = true
	// ch = .: dp[i, j] = true
	//     if dp[i - 1, j - 1] = true
	// ch = *: any i < k < n where s[k] == previousChar: dp[k, j] = true

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(p))
	}
	if !(s[0] == p[0] || p[0] == '.') {
		return false
	}

	dp[0][0] = true
	for sIndex := 1; sIndex < len(s); sIndex++ {
		for pIndex := 1; pIndex < min(sIndex+1, len(p)); pIndex++ {
			if dp[sIndex][pIndex] {
				continue
			}
			if !dp[sIndex-1][pIndex-1] {
				continue
			}

			switch p[pIndex] {
			case '*':
				if pIndex == 0 {
					return false
				}
				previousChar := p[pIndex-1]
				for i := sIndex; i < len(s); i++ {
					if previousChar == '.' {
						dp[i][pIndex] = true
						continue
					}
					if previousChar != s[i] {
						break
					}
					dp[i][pIndex] = true
				}
			case '.':
				dp[sIndex][pIndex] = true
			default:
				dp[sIndex][pIndex] = s[sIndex] == p[pIndex]
			}
		}
	}
	// fmt.Printf("%#v\n", dp)

	return dp[len(s)-1][len(p)-1]
}

func matchRegularExpression(s string, p string) bool {
	queue := make([][]int, 0)
	queue = append(queue, []int{
		0, 0,
	})

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		sIndex := element[0]
		pIndex := element[1]
		if sIndex >= len(s) && pIndex >= len(p) {
			return true
		} else if sIndex >= len(s) || pIndex >= len(p) {
			continue
		}

		switch p[pIndex] {
		case '*':
			if pIndex == 0 {
				return false
			}
			previousChar := p[pIndex-1]
			if previousChar == '.' {
				// match any element, check
				var nextChar byte
				nextPIndex := -1
				for i := pIndex; i < len(p); i++ {
					if p[i] != '.' && p[i] != '*' {
						nextChar = p[i]
						nextPIndex = i
						break
					}
				}
				if nextPIndex != -1 {
					for i := sIndex; i < len(s); i++ {
						if s[i] == nextChar {
							queue = append(queue, []int{
								i,
								nextPIndex,
							})
						}
					}
				}
			} else if previousChar != '*' {
				// s = aaaaab
				// p = a*ab or a*.b
				var i int
				for i = sIndex; i < len(s); i++ {
					if s[i] == previousChar {
						queue = append(queue, []int{
							i,
							pIndex + 1,
						})
					}
				}
			}
		case '.':
			queue = append(queue, []int{
				sIndex + 1,
				pIndex + 1,
			})
		default:
			if s[sIndex] == p[pIndex] {
				queue = append(queue, []int{
					sIndex + 1,
					pIndex + 1,
				})
			} else {
				return false
			}
		}
	}

	return false
}
