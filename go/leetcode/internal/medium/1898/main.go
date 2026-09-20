// https://leetcode.com/problems/maximum-number-of-removable-characters/description/
package main

func maximumRemovals(s string, p string, removable []int) int {
	isSubSequence := func(letters []byte, p string) bool {
		letterIndex := 0
		for pIndex := range p {
			pCh := p[pIndex]
			wasFound := false
			for ; letterIndex < len(letters); letterIndex++ {
				if pCh == letters[letterIndex] {
					letterIndex++
					wasFound = true
					break
				}
			}
			if !wasFound {
				return false
			}
		}
		return true
	}

	letters := []byte(s)
	left := 0
	right := len(removable)

	// O(log2(N))
	for left <= right {
		mid := (left + right) / 2

		for i := 0; i < mid; i++ {
			sIndex := removable[i]
			letters[sIndex] = '.'
		}
		// O(N)
		if isSubSequence(letters, p) {
			// we can still remove more
			left = mid + 1
		} else {
			right = mid - 1
			for i := 0; i < mid; i++ {
				sIndex := removable[i]
				letters[sIndex] = s[sIndex]
			}
		}
	}
	return right
}

/*
func maximumRemovals(s string, p string, removable []int) int {
	pChMap := make(map[byte]int, 0)
	for i := range p {
		ch := p[i]
		pChMap[ch]++
	}

	sChIndexes := make(map[byte][]int, 0)
	for sIndex := range s {
		sCh := s[sIndex]
		if _, ok := pChMap[sCh]; !ok {
			// we don't need an index for letters for non subsequences
			continue
		}

		if _, ok := sChIndexes[sCh]; !ok {
			sChIndexes[sCh] = make([]int, 0)
		}
		sChIndexes[sCh] = append(sChIndexes[sCh], sIndex)
	}

    // all combinations from sChIndexes and pChMap
    combinations :=
    for i := range p {
        ch := p[i]

    }

    k := 0
    for ; k < len(removable); k++ {
        removedIndex := removable[k]
    }
    return k
}
*/
